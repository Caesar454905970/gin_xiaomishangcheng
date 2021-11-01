
FROM golang:alpine

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
	GOPROXY="https://goproxy.cn,direct"

MAINTAINER chengqiang
WORKDIR 17_gin_demo17
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY *.go ./
RUN build -o docker_go .
EXPOSE 8080

CMD ["./docker_go"]

