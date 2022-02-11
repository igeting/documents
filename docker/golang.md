# golang

## build golang docker
```
FROM alpine:3.10

RUN echo http://mirrors.aliyun.com/alpine/v3.10/main > /etc/apk/repositories && \
    echo http://mirrors.aliyun.com/alpine/v3.10/community >> /etc/apk/repositories

RUN apk add --no-cache tzdata && \
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    apk del tzdata

RUN apk add --no-cache libc6-compat ca-certificates && \
    echo "hosts: files dns" > /etc/nsswitch.conf
    
ADD go1.16.10.linux-amd64.tar.gz /usr/local

ENV PATH $PATH:/usr/local/go/bin
```