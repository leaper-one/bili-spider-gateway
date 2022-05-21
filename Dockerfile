# 基础镜像
FROM golang:1.17.8-alpine as builder

RUN apk --no-cache add git
MAINTAINER cunoe

# 操作目录
WORKDIR /go/src/ginForBH/

# 复制源文件至操作目录
COPY . .

# 编译
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GIN_MODE=release go build -o app main.go

FROM alpine:latest as prod

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /go/src/ginForBH/app .

CMD ["./app"]