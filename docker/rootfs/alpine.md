# rootfs

## alpine amd
```
FROM scratch
ADD alpine-minirootfs-3.10.0-amd.tar.gz /
CMD ["/bin/sh"]
```