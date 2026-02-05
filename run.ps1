param(
    [ValidateSet('dev','prod')]
    [string]$mode = 'dev'
)

if ($mode -eq 'dev') {
    $env:DEV_TEMPLATES = '1'
    Write-Host 'DEV_TEMPLATES=1 (templates reload on each request)'
} else {
    Remove-Item Env:DEV_TEMPLATES -ErrorAction SilentlyContinue
    Write-Host 'DEV_TEMPLATES cleared (templates cached)'
}

Write-Host 'Starting server...'

if (Test-Path .\go.mod) {
    go mod tidy
}

go run .\main.go