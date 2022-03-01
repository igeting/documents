# gcc

## options

### -o \<file\> output file

### -g contain debug info

### -v print compile info

### -I \<dir\> add include dir

### -L \<dir\> add library dir 

### -static link static libaray

### -shared dynamic build, link dynamic library

### -lxxx link xxx libaray

### -E preprocess only, do not compile, assemble or link
```
gcc -E *.c -o out.i
```

### -S compile only, do not assemble or link
```
gcc -S *.i -o out.s
```

### -c assemble only, do not link
```
gcc -c *.s -o out.o
```

### link
```
gcc *.o -o x
```

## other options

### -Wl,\<options\> pass comma-separated \<options\> on to the linker

- static link
```
-Wl,-Bstatic
```
- dynamic link
```
-Wl,-Bdynamic
```

> example

```
-Wl,-Bstatic -lpgm -lxerces-c -Wl,-Bdynamic -libverbs -lcurl -Wl,--as-needed
```

## static libaray

### single object
```
gcc -c -o x.o x.c
ar -rc libx.a x.o
```

### multiple object
```
gcc -c -o x1.o x1.c
gcc -c -o x2.o x2.c
ar -rc libx.a x1.o x2.o
```

## shared libaray

### single source
```
gcc -shared -fPIC -o libx.so x.c
```

### single object
```
gcc -shared -fPIC -o libx.so x.o
```

### multiple source
```
gcc -shared -fPIC -o libx.so x1.c x2.c
```

### multiple object
```
gcc -shared -fPIC -o libx.so x1.o x2.o
```

## use shared library (-static use static mode)
```
gcc -c -o xxx.o xxx.c
ar -rc libxxx.a xxx.o
gcc -o test -L. -lxxx test.c [-static]
```
> or
```
gcc -o test test.c libxxx.a
```


## use shared library (windows use libxxx.dll)
```
gcc -shared -fPIC -o libxxx.so xxx.c
gcc -o test -L. -lxxx test.c
```
> or
```
gcc -o test test.c libxxx.so
```