FROM golang:1.24-alpine as API
LABEL authors="pato98"

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o main.go

EXPOSE 8080

ENTRYPOINT exec go run main.go