# Other

## 为非root用户添加docker权限

### 1.添加 docker group 
```
sudo groupadd docker
```

### 2.将用户加入该 group 内
```
$ sudo usermod -aG docker $USER
# or
$ sudo gpasswd -a ${USER} docker
```

### 3.重启服务
```
$ sudo service docker restart
# or
$ sudo /etc/init.d/docker restart
```

### 4.切换当前会话到新 group 或者重启 X 会话

注意:这一步是必须的，否则因为 groups 命令获取到的是缓存的组信息，刚添加的组信息未能生效，所以 docker images 执行时同样有错。

```
newgrp - docker
```

### 5.Centos8安装FFMPEG
```
yum install https://download1.rpmfusion.org/free/el/rpmfusion-free-release-8.noarch.rpm
yum install http://rpmfind.net/linux/epel/7/x86_64/Packages/s/SDL2-2.0.10-1.el7.x86_64.rpm
yum install ffmpeg
```
[SDL](../static/file/linux/SDL.rpm)