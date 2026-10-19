# Этап 1: Сборка бинарного файла
FROM golang:1.27-alpine AS builder

WORKDIR /app

# Кэшируем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код
COPY . .

# Собираем оптимизированный бинарник Go
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/server ./cmd/app/main.go

# Этап 2: Финальный минималистичный образ для запуска
FROM alpine:latest

WORKDIR /app

# Копируем скомпилированный бинарник и шаблоны
COPY --from=builder /app/server .
COPY --from=builder /app/web ./web

EXPOSE 8080

CMD ["./server"]