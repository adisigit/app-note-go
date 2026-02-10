FROM golang:1.25.6-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go build -o app-note-go main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/app-note-go .
COPY --from=builder /app/migration ./migration

EXPOSE 3000

CMD ["./app-note-go"]