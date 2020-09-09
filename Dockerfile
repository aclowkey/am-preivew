# builder
FROM golang:1.14-alpine AS builder

RUN set -ex &&\
    apk add --no-progress --no-cache \
      gcc \
      git \
      musl-dev

WORKDIR /app
COPY . .
RUN GOOS=linux GOARCH=amd64 go build -o am-preview

# image
FROM alpine:edge
WORKDIR /
COPY --from=builder /app/am-preview .
COPY web /web
ENTRYPOINT ["/am-preview"]
