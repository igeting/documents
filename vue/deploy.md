# deploy

## caddy

### caddy file
```
:8080 {
    root * /opt/dist/{$RUN_ENV}
    file_server
    encode zstd gzip
    reverse_proxy /api/* {
        to {$SER_URL}
    }
}
```

### docker file
```
FROM docker.io/itesting/alpine:caddy
COPY dist /code/dist
COPY Caddyfile /code/Caddyfile
ENTRYPOINT ["caddy", "run", "--config", "/code/Caddyfile"]
```

```
FROM docker.io/itesting/npm:caddy
COPY . /code
WORKDIR /code
RUN npm install
RUN npm run build -- --dest="dist/stg" --mode="stg"
RUN npm run build -- --dest="dist/prd" --mode="prd"
ENTRYPOINT ["caddy", "run", "--config", "/code/Caddyfile"]
```

### build docker
```
docker build -t vue-demo .
```

### run docker
```
docker run -d -e RUN_ENV=stg -e SER_URL=http://example:8080 -p 8080:8080 vue-demo
```
