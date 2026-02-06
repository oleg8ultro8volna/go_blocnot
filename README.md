# Go Web блокнот

Ниже полный список шагов и команд, чтобы развернуть, запустить и запушить проект с самого начала.

## 1) Перейти в папку проекта

```powershell
cd d:\PROEKTI_WEB\go_test
```

## 2) Установить Go (одной командой)

```powershell
winget install --id GoLang.Go -e
```

Проверка установки:

```powershell
go version
```

## 3) Запуск проекта

```powershell
go mod tidy
go run main.go
```

Открыть в браузере:

```
http://localhost:8080
```

## 4) Если порт 8080 занят

Проверить, кто занял порт:

```powershell
netstat -ano | findstr :8080
```

Посмотреть процесс:

```powershell
tasklist /FI "PID eq 4876"
```

Остановить процесс (если нужно):

```powershell
taskkill /PID 4876 /F
```

## 5) Перечитывание шаблонов (DEV режим)

Разовый запуск:

```powershell
$env:DEV_TEMPLATES="1"
go run main.go
```

Или через скрипт:

```powershell
.\run.ps1 dev
```

Обычный режим (кеш):

```powershell
.\run.ps1 prod
```

## 6) Инициализация git

```powershell
git init
git remote add origin https://github.com/oleg8ultro8volna/go_blocnot.git
git add .
git commit -m "Initial commit: Go web notebook"
git branch -M main
```

## 7) Установка GitHub CLI

```powershell
winget install --id GitHub.cli -e
```

Проверка:

```powershell
gh --version
```

## 8) Авторизация в GitHub

```powershell
gh auth login
```

Выбрать:
- GitHub.com
- HTTPS
- Yes
- Login with a web browser

Ввести одноразовый код в браузере и подтвердить вход.

## 9) Пуш в репозиторий

```powershell
cd d:\PROEKTI_WEB\go_test
git push -u origin main
```

## 10) Как пользоваться Git и GitHub дальше

Проверить состояние:

```powershell
git status -sb
```

Посмотреть историю:

```powershell
git log --oneline --decorate --graph -n 10
```

### Коммиты

1. Добавить изменения в индекс:
```powershell
git add .
```

2. Создать коммит:
```powershell
git commit -m "Короткое описание изменения"
```

3. Отправить на GitHub:
```powershell
git push
```

### Ветки

Создать новую ветку и перейти в нее:

```powershell
git checkout -b feature/имя-ветки
```

Посмотреть список веток:

```powershell
git branch
```

Переключиться на ветку:

```powershell
git checkout имя-ветки
```

Слить ветку в main:

```powershell
git checkout main
git merge feature/имя-ветки
git push
```

### Вернуться назад

Отменить изменения в конкретном файле (вернуть как было в последнем коммите):

```powershell
git checkout -- path\to\file
```

Отменить все изменения в рабочих файлах (осторожно):

```powershell
git checkout -- .
```

Отменить последний коммит, но оставить изменения в файлах:

```powershell
git reset --soft HEAD~1
```

Полностью удалить последний коммит и изменения (осторожно):

```powershell
git reset --hard HEAD~1
```

### Получить обновления с GitHub

```powershell
git pull
```