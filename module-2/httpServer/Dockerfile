FROM golang:1.18-alpine AS build

#新建工作目录
RUN mkdir -p /httpServer
WORKDIR /httpServer
COPY . /httpServer

#编译httpServer
ENV GO111MODULE=on
RUN go mod download
RUN go build httpServer.go


#生产最终的镜像，仅打包httpServer可执行文件
FROM alpine
MAINTAINER hufeifei
#将编译后的可执行文件COPY到当前镜像
COPY --from=build /httpServer/httpServer /httpServer/httpServer
#启动命令
ENTRYPOINT ["/httpServer/httpServer"]
#定义发布端口
EXPOSE 80