FROM golang:1.17

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
	GOPROXY="https://goproxy.cn,direct"
	
MAINTAINER "chengqiang3"
 
WORKDIR /home/18_gin_demo18
 
ADD . /home/18_gin_demo18
 
RUN go build main.go
 
EXPOSE 8081
 
ENTRYPOINT ["./main"]