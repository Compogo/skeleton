# Многостадийная сборка для минимизации размера образа.
#
# Стадия 1: builder — собирает приложение.
#   - Использует официальный образ Go.
#   - Скачивает зависимости.
#   - Запускает тесты.
#   - Собирает бинарник.
#   - Копирует шаблоны (template/) и миграции (migrations/) опционально.
#
# Стадия 2: alpine — финальный образ.
#   - Использует минимальный Alpine Linux.
#   - Устанавливает сертификаты и часовой пояс.
#   - Копирует бинарник из builder.
#   - Копирует шаблоны и миграции (если они есть).
#
# Результат: небольшой (~20-30 МБ) образ с приложением.

FROM golang:1.26-alpine3.23 as builder

WORKDIR /app

# Копируем go.mod и go.sum для кэширования зависимостей
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код и собираем приложение
COPY . .
RUN GOOS=linux go test --timeout=10m ./...
RUN GOOS=linux go build -a -o ./app ./main.go

# Копируем шаблоны (если есть)
RUN mkdir -p /app/optional_template
RUN if [ -d "./template" ]; then cp -r ./template/. /app/optional_template/; fi

# Копируем миграции (если есть)
RUN mkdir -p /app/optional_migrations
RUN if [ -d "./migrations" ]; then cp -r ./migrations/. /app/optional_migrations/; fi

# Финальный образ на Alpine
FROM alpine

# Устанавливаем сертификаты и часовой пояс
RUN apk add --no-cache ca-certificates tzdata

ENV TZ=Europe/Moscow
WORKDIR /app

# Копируем бинарник и опциональные файлы
COPY --from=builder /app/app /app/app
COPY --from=builder /app/optional_template/. /app/template
COPY --from=builder /app/optional_migrations/. /app/migrations

# Точка входа
ENTRYPOINT ["/app/app"]

# Аргументы командной строки (можно переопределить при запуске)
CMD [""]