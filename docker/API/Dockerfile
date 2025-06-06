FROM golang:1.24-alpine as API

WORKDIR /app

COPY ../app /app

RUN go mod download

RUN go run main.go