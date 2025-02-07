# 使用官方 Go 1.23 镜像作为构建环境
FROM golang:1.23-alpine AS builder

# 设置工作目录
WORKDIR /app

# 将 Go 模块相关文件复制到工作目录
COPY go.mod go.sum ./

# 下载 Go 依赖
RUN go mod tidy

# 复制 server 和 client 代码
COPY server /app/server
COPY client /app/client

# 编译服务端和客户端
RUN go build -o /app/server/server /app/server/main.go
RUN go build -o /app/client/client /app/client/main.go

# 使用更小的镜像来运行应用
FROM alpine:latest

# 设置工作目录
WORKDIR /app

# 将编译后的二进制文件复制到运行镜像
COPY --from=builder /app/server/server /app/server
COPY --from=builder /app/client/client /app/client

# 开放端口
EXPOSE 8080

# 默认运行服务端
CMD ["/app/server"]
# 使用官方 Go 1.23 镜像作为构建环境
FROM golang:1.23-alpine AS builder

# 设置工作目录
WORKDIR /app

# 将 Go 模块相关文件复制到工作目录
COPY . .

# 下载 Go 依赖
RUN go mod tidy


# 编译服务端和客户端
RUN go build -o /app/server/server /app/server/main.go
RUN go build -o /app/client/client /app/client/main.go

# 使用更小的镜像来运行应用
FROM alpine:latest

# 设置工作目录
WORKDIR /app

# 将编译后的二进制文件复制到运行镜像
COPY --from=builder /app/server/server /app/server
COPY --from=builder /app/client/client /app/client

# 开放端口
EXPOSE 8080

# 默认运行服务端
CMD ["/app/server"]
