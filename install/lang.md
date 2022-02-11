# install language

## ubuntu

### search language
```
apt search language-pack
```

### install zh_CN
```
apt install language-pack-zh-hans
```

### install en_US
```
apt install language-pack-en
```

### set language
```
vi /etc/profile (~/.bashrc)
export LANG=en_US.UTF-8
export LANGUAGE=en_US.UTF-8
export LC_ALL=en_US.UTF-8
```

## centos

### search language
```
yum search langpacks
```

### install zh_CN
```
yum install glibc-all-langpacks
yum install langpacks-zh_CN
```

### install en_US
```
yum install glibc-all-langpacks
yum install langpacks-en_US
```