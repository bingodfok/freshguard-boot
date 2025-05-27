FROM golang:1.24.0 AS builder

ENV GOOS=linux \
    GOARCH=amd64 \
    CGO_ENABLED=0 \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /builds

COPY . .

RUN go mod tidy
RUN go build -o app ./cmd/main.go


FROM alpine:latest

EXPOSE 8080

WORKDIR /app

COPY --from=builder /builds/app .
COPY --from=builder /builds/cmd/etc ./cmd/etc

RUN apk add dumb-init
ENTRYPOINT ["/usr/bin/dumb-init", "--"]

CMD ["./app"]