API‑прокси для ML‑моделей (hba1c, ldl, tg, ferr, hdl)

Принимает GET‑запросы с параметрами в URL, перенаправляет их на внешний API, обрабатывает ответы и возвращает JSON.

---

# Структура проекта

awesomeProject/
├── cmd/
│ └── app/
│ └── main.go # точка входа, загрузка конфига, настройка маршрутов
├── config/
│ ├── config.go # логика парсинга config.yml
│ └── config.yml # пример файла конфигурации
├── internal/
│ ├── handler/ # HTTP‑хендлеры
│ │ ├── all.go # объединённый эндпоинт /predict/all
│ │ ├── hba1c.go
│ │ ├── ldl.go
│ │ ├── tg.go
│ │ ├── ferr.go
│ │ └── hdl.go
│ ├── middleware/ # авторизация
│ │ └── auth.go
│ ├── models/ # структуры запросов/ответов
│ │ ├── all.go
│ │ ├── hba1c.go
│ │ ├── ldl.go
│ │ ├── tg.go
│ │ ├── ferr.go
│ │ └── hdl.go
│ └── services/ # бизнес‑логика: сборка и отправка запросов во внешний API
│ ├── all.go
│ ├── hba1c.go
│ ├── ldl.go
│ ├── tg.go
│ ├── ferr.go
│ ├── hdl.go
│ └── request.go
├── build.sh # Linux‑скрипт сборки
├── build.bat # Windows‑скрипт сборки
├── Dockerfile # сборка Docker‑образа
├── docker‑compose.yml # поднятие контейнера со всеми зависимостями
├── go.mod # зависимости Go-модуля
└── README.md # этот файл



---

# Конфигурация

В корневой папке лежит `config/config.yml`:

port: 8080
auth_token: "0l62<EJi/zJx]a?"
api_urls:
  hba1c: "https://apiml.labhub.online/api/v1/predict/hba1c"
  ldl:   "https://apiml.labhub.online/api/v1/predict/ldl"
  tg:    "https://apiml.labhub.online/api/v1/predict/tg"
  ferr:  "https://apiml.labhub.online/api/v1/predict/ferr"
  hdl:   "https://apiml.labhub.online/api/v1/predict/hdl"





# Локальная сборка и запуск
cd /path/to/awesomeProject
go mod tidy
go build -o awesomeProject.exe ./cmd/app
# Windows
.\awesomeProject.exe
# Linux/macOS
./awesomeProject.exe






# Docker & docker‑compose

Dockerfile:
FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
WORKDIR /app/cmd/app
RUN go build -o /awesomeProject

FROM alpine:3.18
WORKDIR /app
COPY --from=builder /awesomeProject .
COPY config.yml ./config.yml
COPY config/ ./config/
EXPOSE 8080
ENTRYPOINT ["./awesomeProject"]

docker-compose.yml:
version: '3.8'
services:
  app:
    build: .
    ports:
      - "8080:8080"
    volumes:
      - ./config.yml:/app/config.yml:ro
      - ./config/:/app/config/:ro
Запуск:
docker-compose up --build -d





# Эндпоинты
Все запросы требуют заголовков:
Authorization: Bearer <auth_token>
Accept: application/json; charset=utf-8

GET /predict/hba1c
GET /predict/ldl
GET /predict/tg
GET /predict/ferr
GET /predict/hdl
GET /predict/all

пример curl:
curl -G 'http://localhost:8080/predict/hba1c' \
  --data-urlencode 'uid=web-client' \
  …(остальные параметры)… \
  -H 'Authorization: Bearer 0l62<EJi/zJx]a?' \
  -H 'Accept: application/json; charset=utf-8'



# Тестирование
Postman: создайте GET‑запрос, вставьте URL, в Headers добавьте Authorization и Accept, нажмите Send.


# Contributing
Создайте репозиторий
Создайте ветку feature/ваша-ветка
Внесите изменения, добавьте тесты
Откройте Pull Request в main
