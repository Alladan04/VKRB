# Используем официальный образ Go как базовый для сборки
FROM golang:1.23.4 AS builder

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем go.mod и go.sum и устанавливаем зависимости
COPY ./MivarAPI/go.mod ./MivarAPI/go.sum ./
RUN go mod download

# Копируем остальной исходный код
COPY ./MivarAPI .

# Собираем Go-приложение
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server ./cmd/service

# Используем минимальный образ для запуска
FROM alpine:latest

# Устанавливаем рабочую директорию в финальном образе
WORKDIR /app

# Копируем бинарник из стадии сборки
COPY --from=builder /app/server .

# Копируем конфиг и папку data
COPY ./MivarAPI/config.yaml .
COPY ./MivarAPI/data ./data

# Указываем порт (при необходимости)
EXPOSE 8080

# Команда запуска
CMD ["./server"]