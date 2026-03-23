FROM golang:1.26-alpine3.23 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

# build app
COPY . .
RUN GOOS=linux go test --timeout=10m ./...
RUN GOOS=linux go build -a -o ./app ./main.go

# copy template
RUN mkdir -p /app/optional_template
RUN if [ -d "./template" ]; then cp -r ./template/. /app/optional_template/; fi

# copy migrations
RUN mkdir -p /app/optional_migrations
RUN if [ -d "./migrations" ]; then cp -r ./migrations/. /app/optional_migrations/; fi

FROM alpine

RUN apk add --no-cache ca-certificates tzdata

ENV TZ=Europe/Moscow
WORKDIR /app

COPY --from=builder /app/app /app/app
COPY --from=builder /app/optional_template/. /app/template
COPY --from=builder /app/optional_migrations/. /app/migrations

ENTRYPOINT ["/app/app"]

CMD [""]
