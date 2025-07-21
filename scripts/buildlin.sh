#!/bin/bash

mkdir -p release/linux

GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o release/linux/api-gateway ./cmd/app/

cp config.yml release/linux/
cp .env release/linux/

chmod +x release/linux/api-gateway

echo "Сборка завершена. Результат в папке release/"