FROM golang:1.22.2-bookworm

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
# 设置时区为上海
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo 'Asia/Shanghai' >/etc/timezone
 
# 设置时区（以 Asia/Shanghai 为例）
ENV TZ=Asia/Shanghai
 
# 设置编码
ENV LANG C.UTF-8