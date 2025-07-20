# syntax=docker/dockerfile:1

FROM golang:1.23.5 as build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o ./main main.go

# ---- RUNTIME ----
FROM debian:bullseye-slim

WORKDIR /app

COPY --from=builder /app/main ./main
# COPY .env .env

CMD ["./main"]
