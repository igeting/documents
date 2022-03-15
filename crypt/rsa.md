# openssl rsa

## generate private key
```
openssl genrsa -out private.pem 1024
```

## parse PKCS (not necessary)
```
openssl pkcs8 -topk8 -inform PEM -in private.pem -outform PEM –nocrypt
```

## generate public key
```
openssl rsa -in private.pem -pubout -out public.pem
```
