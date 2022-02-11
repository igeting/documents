# redis docker

## start a redis instance
```
docker run --name some-redis -d redis:v6.0.0
```

## list the network
```
docker network list
docker network create some-network
```

## start with persistent storage
```
docker run --name some-redis -d redis:v6.0.0 redis-server --appendonly yes
```

## connecting via redis-cli
```
#ser
docker run -d --name some-redis -p 6379:6379 --network some-network redis:v6.0.0 redis-server
#cli
docker run -it --rm --network some-network redis:v6.0.0 redis-cli -h some-redis -p 6379 -a 123456
```

## allow remote access
```
docker run -d --name some-redis -p 6379:6379 -v redis.conf:/redis.conf -v data:/data redis:v6.0.0 redis-server /redis.conf --appendonly yes

#redis.conf
bind 0.0.0.0
daemonize NO
protected-mode no
requirepass 123456
```

## Additionally, If you want to use your own redis.conf ...
```
docker run --name some-redis -v /myredis/conf/redis.conf:/usr/local/etc/redis/redis.conf redis:v6.0.0 redis-server /usr/local/etc/redis/redis.conf
```

## redis set password
```
config get requirepass
config set requirepass 123456
```