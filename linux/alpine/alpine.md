# alpine

## cmd

### source
```
echo http://mirrors.aliyun.com/alpine/v3.10/main > /etc/apk/repositories
echo http://mirrors.aliyun.com/alpine/v3.10/community >> /etc/apk/repositories
```

### update
```
apk update
apk upgrade
```

### info
```
apk info xxx
apk info -a xxx
apk info --who-owns /bin/xxx
```

### add xxx
```
apk add xxx
apk add xxx=version
apk add 'xxx<version'
apk add 'xxx>version'
apk add --upgrade xxx
apk add --no-cache xxx
apk add --update-cache xxx
apk add --repository http://mirrors.aliyun.com/alpine/v3.10/main --allow-untrusted xxx
```

### del xxx
```
apk del xxx && \
rm -rf /var/cache/apk/* && \
rm -rf /root/.cache && \
rm -rf /tmp/*
```

### search xxx
```
apk search xxx
apk search -v xxx
```

## service

### install rc-service
```
apk add --no-cache openrc
```

### list service
```
rc-service --list
```

### manage xxx service
```
rc-service xxx start/stop/restart

#or
/etc/init.d/xxx start/stop/restart
```

### set boot xxx
```
rc-update add xxx
```

### set service (/etc/init.d/xxx)
```
#!/sbin/openrc-run
 
name="actc"
command="/path/to/${name}"
#command_background="yes"
 
depend() {
	after sshd
}
```

## lib

### musl compatible glibc
```
apk add --no-cache libc6-compat

#or amd
mkdir /lib64
ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

#or arm
ln -s /lib/libc.musl-aarch64.so.1 /lib/ld-linux-aarch64.so.1
```

## docker

### use shanghai timezone
```
COPY localtime /etc/localtime
```

### golang
```
FROM scratch
ADD arm.tar.gz /
ADD go.xxx.tar.gz /usr/local
ENV PATH $PATH:/usr/local/go/bin
CMD ["/bin/sh"]
```