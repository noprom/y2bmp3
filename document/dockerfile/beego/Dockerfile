#FROM golang:1.6.0
FROM daocloud.io/library/golang:1.8.1
MAINTAINER noprom tyee.noprom@qq.com

COPY bin/* /usr/local/bin/
# Install beego & bee
RUN go get -u github.com/astaxie/beego
RUN go get -u github.com/beego/bee
RUN curl https://glide.sh/get | sh
