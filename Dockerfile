FROM golang:alpine as builder

WORKDIR /build
COPY . .

ARG GOARCH="amd64"
RUN go get -u ./... && \
    GOOS=linux GOARCH=$GOARCH go build -ldflags="-w -s" -o build

FROM alpine:latest

WORKDIR /build
COPY --from=builder /build/build .

ENV MONKEBASE_CONNECTION ""
ENTRYPOINT ./build
