# syntax=docker/dockerfile:1

FROM golang:1.21 as build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o bot .

# ---- RUNTIME ----
FROM debian:bullseye-slim

WORKDIR /app

COPY --from=build /app/bot .
COPY .env .env

CMD ["./bot"]
