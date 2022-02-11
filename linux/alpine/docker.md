# alpine

## login docker
```
docker login docker.io
```

## build docker
```
FROM scratch
ADD rootfs.tar.gz /
CMD ["/bin/sh"]
```

### change source
```
FROM alpine:3.10
RUN \
    echo http://mirrors.aliyun.com/alpine/v3.10/main > /etc/apk/repositories && \
    echo http://mirrors.aliyun.com/alpine/v3.10/community >> /etc/apk/repositories
```

### change timezone
```
FROM alpine:3.10
RUN apk add --no-cache tzdata && \
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    apk del tzdata
```

### build golang env
```
FROM alpine:3.10
RUN apk add --no-cache libc6-compat ca-certificates && \
    echo "hosts: files dns" > /etc/nsswitch.conf
ADD go.xxx.tar.gz /usr/local
ENV PATH $PATH:/usr/local/go/bin
```