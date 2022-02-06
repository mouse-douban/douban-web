# 使用官方 Golang 镜像作为构建环境
FROM golang:1.17-buster as builder

WORKDIR /app

# 换成阿里云镜像
RUN sed -i "s/archive.ubuntu./mirrors.aliyun./g" /etc/apt/sources.list
RUN sed -i "s/deb.debian.org/mirrors.aliyun.com/g" /etc/apt/sources.list
RUN sed -i "s/security.debian.org/mirrors.aliyun.com\/debian-security/g" /etc/apt/sources.list

ENV BUCKET_URL=xxx
ENV TENCENT_SECRET_ID=xxx
ENV TENCENT_SECRET_KEY=xxx

# 安装依赖
COPY go.* ./
# 设置代理
RUN go env -w GOPROXY="https://goproxy.cn,direct"
# 下载依赖
RUN go mod tidy

# 将代码文件写入镜像
COPY . ./

# 使用远程config
RUN sed -i 's/EnableCOS = false/EnableCOS = true/g' cmd/main.go

# 构建二进制文件
RUN go build -mod=readonly -v -o server cmd/main.go
# 尝试下载 tls key
RUN go run build/config.go

# 使用裁剪后的官方 Debian 镜像作为基础镜像
# https://hub.docker.com/_/debian
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM debian:buster-slim
RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*

# 将构建好的文件拷贝进镜像
COPY --from=builder /app/server /app/server
COPY --from=builder /app/config/api.key /app/config/api.key
COPY --from=builder /app/config/api.pem /app/config/api.pem

# 启动 Web 服务
CMD ["/app/server"]
