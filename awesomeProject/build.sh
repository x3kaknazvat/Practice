#!/usr/bin/env bash
set -euo pipefail

# Собираем бинарь в корне проекта
go build -o awesomeProject ./cmd/app
echo "Собран awesomeProject"
