package main

/*
#cgo CFLAGS: -I./
#cgo LDFLAGS: -L./ -ltest
#include "test.h"
*/
import "C"
import "fmt"

//static
/*
gcc -c test.c
ar -cr libtest.a test.o
go build -o app test.go
*/

//share
/*
gcc -shared -fPIC -o libtest.so test.c
go build -ldflags="-r ./" -o app test.go
*/
func main() {
	s1 := C.test(C.CString("golang"))
	fmt.Println(C.GoString(s1))
	s2 := C.test(C.CString(""))
	println(C.GoString(s2))
}
