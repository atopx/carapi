package model

const BuildDockerFile = `FROM golang:alpine as builder

# set build workdir
WORKDIR /build

RUN apk add tzdata

# set env
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# 复制项目中的 go.mod 和 go.sum文件并下载依赖信息
COPY go.mod .
COPY go.sum .
RUN go mod download

# copy source code
COPY . .

# run build
RUN go build -ldflags="-s -w" -tags=jsoniter -o app

# release
FROM alpine:latest
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
COPY --from=builder /build/app /app
`
const LocalDockerFile = `FROM golang:alpine

RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

COPY go.mod .
COPY go.sum .
RUN go mod download
`
