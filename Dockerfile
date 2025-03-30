FROM golang:1.24.1 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY .env.docker ./.env

COPY *.go ./

COPY config config
COPY database database
COPY handler handler
COPY middleware middleware
COPY model model
COPY router router
COPY utils utils

RUN CGO_ENABLED=0 GOOS=linux go build -o /companies-microservice

EXPOSE 8080

ENTRYPOINT ["/companies-microservice"]
