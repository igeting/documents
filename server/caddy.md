# caddy

## config file (Caddyfile)
```
:9090 {
    root * /opt/dist
    file_server
    reverse_proxy /api/* 127.0.0.1:8080
}
```

## vue dist

### caddy 1
```
:9090 {
    root * /opt/dist
    file_server
    rewrite {
        regexp .*
        to {path} /
    }
}

```

### caddy 2
```
:9090 {
    root * /opt/dist
    file_server
    encode zstd gzip
    try_files {path} /index.html
    log {
        output file /tmp/caddy_log
        format single_field common_log
    }
}
```

## reverse proxy

### vue proxy
```
:9090 {
    root * /opt/dist
    reverse_proxy /api/* {
        to http://localhost:8080
    }
}
```