# Download the image for the specific the platform and named it as builder
FROM --platform=$BUILDPLATFORM golang:1.20.2-alpine3.17 AS builder
WORKDIR /build

ENV GOOS=linux
ENV GOPATH /go
ENV GOCACHE /go-build

COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod/cache \
    go mod download
COPY . .

RUN --mount=type=cache,target=/go/pkg/mod/cache \
    --mount=type=cache,target=/go-build \
    go build -o finance-api ./cmd/api

CMD ["/build/finance-api"]