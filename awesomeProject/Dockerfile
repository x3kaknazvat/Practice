
# === 1. Собираем бинарник ===
FROM golang:1.20-alpine AS builder

WORKDIR /app
# Сначала подтягиваем модули
COPY go.mod go.sum ./
RUN go mod download

# Копируем весь код
COPY . .

# Собираем приложение
WORKDIR /app/cmd/app
RUN go build -o /awesomeProject

# === 2. Собираем финальный образ ===
FROM alpine:3.18

WORKDIR /app

# Копируем бинарь
COPY --from=builder /awesomeProject .

# Копируем конфиг-файл и пакет config/
COPY config.yml .
COPY config/ ./config

EXPOSE 8080

# Запускаем
ENTRYPOINT ["./awesomeProject"]
