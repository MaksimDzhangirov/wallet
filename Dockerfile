# Build stage
FROM golang:1.16-alpine3.13 AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main main.go
RUN apk --no-cache add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz

# Run stage
FROM alpine:3.13
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/migrate.linux-amd64 ./migrate
COPY app.env .
COPY db/migration ./migration

EXPOSE 8080
CMD [ "/app/main" ]