# error

## rpmdb open failed
```
rm -f /var/lib/rpm/__db*
rpm --rebuilddb
yum -y update
```

## cmake (undefined symbol archive_write_add_filter_zstd)
```
yum install libarchive
```