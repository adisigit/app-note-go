FROM golang:1.25.6-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go build -o note-go main.go

CMD ["./note-go"]