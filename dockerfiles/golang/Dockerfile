FROM golang:1.20.2

RUN go install github.com/cespare/reflex@latest

WORKDIR /app
COPY ./api ./

RUN go mod download