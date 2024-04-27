# 使用官方 Golang 镜像作为构建环境
FROM golang:latest AS builder


ENV GOPROXY https://goproxy.cn,direct
# 设置工作目录
WORKDIR /app

# 复制 go.mod 和 go.sum 文件并下载依赖项
COPY go.mod go.sum ./
RUN go mod download

# 复制其他文件
COPY . .

# 构建应用程序
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o myapp .

# 使用 alpine 镜像来获得一个更小的镜像
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# 复制构建器阶段的可执行文件
COPY --from=builder /app/myapp .

# 设置运行应用程序的命令
CMD ["./myapp"]