package main

// #cgo CFLAGS: -g -Wall
// #include <stdlib.h>
// #include <stdio.h>
import "C"
import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	//Seed(uint64(time.Now().UnixMilli()))
	//Random()
	Print("ciao")
}

func Random() int64 {
	var r C.long = C.random()
	fmt.Println("type: ", reflect.ValueOf(r).Type())
	fmt.Println("kind: ", reflect.ValueOf(r).Kind())
	return int64(r)
}

func Seed(i uint64) {
	C.srandom(C.uint(i))
}

// Unlike Go, C doesn’t have an explicit string type. Strings in C are represented by
// a zero-terminated array of chars. Conversion between Go and C strings is done with
// the C.CString, C.GoString, and C.GoStringN functions. These conversions make a copy
// of the string data.

func Print(s string) {
	// from Go string to a C allocated string (allocated on C memory)
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	C.fputs(cs, (*C.FILE)(C.stdout))
}
