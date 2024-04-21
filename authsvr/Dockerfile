# 使用Go官方镜像作为构建环境
FROM golang:1.16 AS builder
WORKDIR /app
COPY . .

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct
RUN go mod download

# 编译Go程序，禁用cgo，并设置目标操作系统为linux
RUN CGO_ENABLED=0 GOOS=linux go build -o authsvr .

# 使用Alpine Linux作为最终镜像，以减少镜像大小
FROM alpine:latest
WORKDIR /root/
# 从构建阶段复制编译好的程序到最终镜像中
COPY --from=builder /app/authsvr .
# 设置容器启动时执行的命令
CMD ["./authsvr"]
