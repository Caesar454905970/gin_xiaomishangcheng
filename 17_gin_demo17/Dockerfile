FROM golang:1.17-alpine
MAINTAINER chengqiang
WORKDIR /go/src/docker_go
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...

CMD ["docker_go"]