# centos

## Synchronize cache AppStream

### CentOS-AppStream.repo
```
[AppStream]
name=CentOS-$releasever - AppStream
baseurl=https://mirrors.aliyun.com/centos/8-stream/AppStream/$basearch/os/
gpgcheck=1
enabled=1
gpgkey=file:///etc/pki/rpm-gpg/RPM-GPG-KEY-centosofficial
```

### CentOS-Base.repo
```
[BaseOS]
name=CentOS-$releasever - Base
baseurl=https://mirrors.aliyun.com/centos/8-stream/BaseOS/$basearch/os/
gpgcheck=1
enabled=1
gpgkey=file:///etc/pki/rpm-gpg/RPM-GPG-KEY-centosofficial
```

### CentOS-Extras.repo
```
[extras]
name=CentOS-$releasever - Extras
baseurl=https://mirrors.aliyun.com/centos/8-stream/extras/$basearch/os/
gpgcheck=1
enabled=1
gpgkey=file:///etc/pki/rpm-gpg/RPM-GPG-KEY-centosofficial
```

### CentOS-PowerTools.repo
```
[PowerTools]
name=CentOS-$releasever - PowerTools
baseurl=https://mirrors.aliyun.com/centos/8-stream/PowerTools/$basearch/os/
gpgcheck=1
enabled=1
gpgkey=file:///etc/pki/rpm-gpg/RPM-GPG-KEY-centosofficial
```
