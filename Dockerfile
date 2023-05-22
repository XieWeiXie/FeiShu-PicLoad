FROM golang:1.20.0 as builder

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    CGO_ENABLED=1 \
    TZ=Asia/Shanghai

ARG APP=FeiShuPicLoad

WORKDIR /app

COPY go.mod .
COPY go.sum .
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /app/target/${APP} main.go && go clean --cache && go clean --modcache

FROM alpine:latest

RUN apk update --no-cache && apk add --no-cache ca-certificates tzdata bash

WORKDIR /app
COPY --from=builder /app/target/${APP} /app/${APP}
EXPOSE 9091
CMD ["/app/FeiShuPicLoad"]
