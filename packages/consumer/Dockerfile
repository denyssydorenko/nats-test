# syntax=docker/dockerfile:1.4
FROM golang:1.19-buster as build-env

RUN mkdir /app
WORKDIR /app

COPY --link ./ .
COPY --link go.mod .
COPY --link go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build ./cmd/consumer/

# Deploy into alpine
FROM alpine:latest

WORKDIR /app
COPY --from=build-env --link /app/consumer .

EXPOSE 8085
CMD ["./consumer"]