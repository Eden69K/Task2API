New-Item -ItemType Directory -Path "release\windows" -Force

$env:GOOS="windows"; $env:GOARCH="amd64"; go build -ldflags="-s -w" -o "release\windows\api-gateway.exe" ./cmd/app/

Copy-Item "config.yml" -Destination "release\windows\"
Copy-Item ".env" -Destination "release\windows\.env"

Write-Host "Сборка завершена. Файлы в папке release"