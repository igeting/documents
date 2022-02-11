## vlmcsd is a replacement for Microsoft's KMS server.
It contains vlmcs, a KMS test client, mainly for debugging purposes, that also can "charge" a genuine KMS server designed to run on an always-on or often-on device, e.g. router, NAS Box, ...intended to help people who lost activation of their legally-owned licenses, e.g. due to a change of hardware (motherboard, CPU, ...)
vlmcsd is not a one-click activation or crack tool intended to activate illegal copies of software (Windows, Office, Project, Visio)

## Info / About this docker
Docker based in Alpine OS with vlmcsd compiled from "source" (vlmcsd GitHub)

## Get Image
```
docker pull ideving/vlmcsd:1.0.0
```

## Server Usage
```
docker run -d -p 1688:1688 --restart=always --name vlmcsd ideving/vlmcsd
```
## To view docker log
Now vlmcsd process send logs to docker.
```
 docker logs vlmcsd (change 'vlmcsd' with the docker's name)
```

## Client
- Windows
```
slmgr.vbs -upk
slmgr.vbs -ipk XXXXX-XXXXX-XXXXX-XXXXX-XXXXX
slmgr.vbs -skms DOCKER_IP
slmgr.vbs -ato
slmgr.vbs -dlv
```
- Office x32
```
cd \Program Files (x86)\Microsoft Office\Office16
cscript ospp.vbs /sethst:DOCKER_IP
cscript ospp.vbs /inpkey:xxxxx-xxxxx-xxxxx-xxxxx-xxxxx
cscript ospp.vbs /act
cscript ospp.vbs /dstatusall
```
- Office x64
```
cd \Program Files\Microsoft Office\Office16
cscript ospp.vbs /sethst:DOCKER_IP
cscript ospp.vbs /inpkey:xxxxx-xxxxx-xxxxx-xxxxx-xxxxx
cscript ospp.vbs /act
cscript ospp.vbs /dstatusall
```

## Docker Link
```
https://hub.docker.com/r/mikolatero/vlmcsd/
```