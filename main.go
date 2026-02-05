package main

import (
    "database/sql"
    "html/template"
    "log"
    "net/http"
    "os"
    "path/filepath"
    "strconv"
    "time"

    _ "modernc.org/sqlite"
)

type Note struct {
    ID        int64
    Content   string
    CreatedAt time.Time
}

type PageData struct {
    Notes []Note
}

func main() {
    db, err := sql.Open("sqlite", "notes.db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    if err := initDB(db); err != nil {
        log.Fatal(err)
    }

    tmplPath := filepath.Join("templates", "index.html")
    devTemplates := os.Getenv("DEV_TEMPLATES") == "1"

    var tmpl *template.Template
    if !devTemplates {
        tmpl, err = template.ParseFiles(tmplPath)
        if err != nil {
            log.Fatal(err)
        }
    }

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodGet {
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
            return
        }

        notes, err := listNotes(db)
        if err != nil {
            http.Error(w, "Failed to load notes", http.StatusInternalServerError)
            return
        }

        t := tmpl
        if devTemplates {
            var err error
            t, err = template.ParseFiles(tmplPath)
            if err != nil {
                http.Error(w, "Failed to load template", http.StatusInternalServerError)
                return
            }
        }

        if err := t.Execute(w, PageData{Notes: notes}); err != nil {
            http.Error(w, "Failed to render", http.StatusInternalServerError)
            return
        }
    })

    http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
            return
        }

        if err := r.ParseForm(); err != nil {
            http.Error(w, "Bad request", http.StatusBadRequest)
            return
        }

        content := r.FormValue("content")
        if len(content) == 0 {
            http.Redirect(w, r, "/", http.StatusSeeOther)
            return
        }

        if err := addNote(db, content); err != nil {
            http.Error(w, "Failed to add note", http.StatusInternalServerError)
            return
        }

        http.Redirect(w, r, "/", http.StatusSeeOther)
    })

    http.HandleFunc("/delete", func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
            return
        }

        if err := r.ParseForm(); err != nil {
            http.Error(w, "Bad request", http.StatusBadRequest)
            return
        }

        idStr := r.FormValue("id")
        id, err := strconv.ParseInt(idStr, 10, 64)
        if err != nil {
            http.Error(w, "Bad request", http.StatusBadRequest)
            return
        }

        if err := deleteNote(db, id); err != nil {
            http.Error(w, "Failed to delete note", http.StatusInternalServerError)
            return
        }

        http.Redirect(w, r, "/", http.StatusSeeOther)
    })

    log.Println("Listening on http://localhost:8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}

func initDB(db *sql.DB) error {
    _, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS notes (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            content TEXT NOT NULL,
            created_at INTEGER NOT NULL
        );
    `)
    return err
}

func addNote(db *sql.DB, content string) error {
    _, err := db.Exec(
        "INSERT INTO notes (content, created_at) VALUES (?, ?)",
        content,
        time.Now().Unix(),
    )
    return err
}

func deleteNote(db *sql.DB, id int64) error {
    _, err := db.Exec("DELETE FROM notes WHERE id = ?", id)
    return err
}

func listNotes(db *sql.DB) ([]Note, error) {
    rows, err := db.Query("SELECT id, content, created_at FROM notes ORDER BY id DESC")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var notes []Note
    for rows.Next() {
        var n Note
        var created int64
        if err := rows.Scan(&n.ID, &n.Content, &created); err != nil {
            return nil, err
        }
        n.CreatedAt = time.Unix(created, 0)
        notes = append(notes, n)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return notes, nil
}
