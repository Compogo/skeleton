# Compogo Skeleton

Скелетон (шаблон) для быстрого старта нового микросервиса на фреймворке [Compogo](https://github.com/Compogo/compogo).

## Структура

```text
sceleton/
├── main.go                      # Точка входа
├── Dockerfile                   # Многостадийная сборка
├── go.mod                       # Зависимости
├── infrastructure/
│   └── config/                  # Конфигурация приложения
│       ├── config.go            # Параметры конфигурации
│       └── component.go         # Компонент для Compogo
├── interface/
│   └── cli/
│       └── root.go              # Корневая команда CLI
├── migrations/                  # Миграции БД (опционально)
└── template/                    # Шаблоны (опционально)
```

## Быстрый старт

```shell
# Клонирование
git clone https://github.com/Compogo/sceleton.git my-service
cd my-service

# Сборка
go build -o app ./main.go

# Запуск
./app
```

## Docker

### Сборка образа

```shell
docker build -t my-service .
```

### Запуск контейнера

```shell
docker run --rm my-service --app.test=hello
```

### Многостадийная сборка

Dockerfile использует многостадийную сборку для минимизации размера образа:

1. builder — собирает приложение на golang:1.26-alpine3.23
2. alpine — финальный образ с минимальным набором (ca-certificates, tzdata)


## Миграции

Если в проекте есть папка migrations/, она копируется в образ и доступна для мигратора.

```shell
# Структура миграций
migrations/
├── mysql/
│   ├── 1_init.up.sql
│   └── 1_init.down.sql
└── postgres/
    ├── 1_init.up.sql
    └── 1_init.down.sql
```

## Шаблоны

Если в проекте есть папка template/, она копируется в образ и доступна для приложения.

## Лицензия

```text
MIT License

Copyright (c) 2026 Compogo

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

```
