package main

/*
#cgo CFLAGS: -I./include
#cgo LDFLAGS: -L./lib -ltest -Wl,-rpath,lib
#include "test.h"
*/
import "C"
import "fmt"

//static
/*
gcc -c test.c
ar -cr ./lib/libtest.a test.o
go build -o app test.go
*/

//share
/*
gcc -shared -fPIC -o ./lib/libtest.so test.c
go build -ldflags="-r ./lib" -o app test.go
*/
func main() {
	s1 := C.test(C.CString("golang"))
	fmt.Println(C.GoString(s1))
	s2 := C.test(C.CString(""))
	println(C.GoString(s2))
}
