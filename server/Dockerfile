# syntax=docker/dockerfile:1

FROM golang:1.19-alpine AS builder
RUN apk add build-base

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . ./

RUN go build -o /bin/server -tags musl main.go

FROM alpine:3.14
COPY --from=builder /bin/server /bin/server

EXPOSE 8080

CMD ["/bin/server"]