# syntax = docker/dockerfile:1-experimental

# Builder
FROM golang:1.19.4-alpine3.17 as builder

RUN apk update && \
    apk upgrade && \
    apk --update add git make bash build-base

WORKDIR /app

COPY go.mod go.sum .

RUN go mod download

COPY . .

RUN --mount=type=cache,target=${GOCACHE} \
    make build

# Distribution
FROM alpine:latest

RUN apk update && \
    apk upgrade && \
    apk --update --no-cache add tzdata && \
    mkdir /app

WORKDIR /app

EXPOSE 9090

COPY --from=builder /app /app/

CMD /app/main
