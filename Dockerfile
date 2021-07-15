FROM golang:alpine as builder

WORKDIR /build
COPY . .

ARG GOARCH="amd64"
ARG GOOS="linux"
RUN go get -u ./... && \
    GOOS=$GOOS GOARCH=$GOARCH go build -ldflags="-w -s" -o build

FROM alpine:latest

WORKDIR /build
COPY --from=builder /build/build .

ENV DATABASE_CONNECTION ""
ENTRYPOINT ./build
