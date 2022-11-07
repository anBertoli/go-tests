package main

/*
#cgo CFLAGS: -g -Wall -I/usr/local/include
#include <stdlib.h>
#include "./c/hello.c"

void myprint(char* s) {
	printf("%s\n", s);
}
*/
import "C"

import (
	"fmt"
	"unsafe"
)

func main() {
	var year C.int = C.int(2018)

	// From Go string to a C allocated string (allocated on C memory). These
	// are dynamically allocates C strings and must be freed after use.
	var name *C.char = C.CString("Gopher")
	defer C.free(unsafe.Pointer(name))

	// The function C.malloc returns an object of type unsafe.Pointer.
	// The space is allocated in C managed memory and must be freed.
	var ptr unsafe.Pointer = C.malloc(C.sizeof_char * 1024)
	defer C.free(ptr)

	// Call the function we defined in the hello.c file. Note the arguments
	// match the "translated" function signature.
	var size C.int = C.hello(name, year, (*C.char)(ptr))
	fmt.Println(size)

	// We copy/convert the C buffer to a go []byte object. The cgo function
	// GoBytes does this for us, using the pointer and the size of the written
	// data. The byte slice returned does not share memory with the bytes we
	// allocated using malloc (it's a copy).
	b := C.GoBytes(ptr, size)
	fmt.Println(string(b))

	// We should check for nil pointers being returned.
	fmt.Println("\n===============")
	var cStr *C.char = C.hello_alloc(name, year)
	defer C.free(unsafe.Pointer(cStr))
	fmt.Printf("C string (%p): ", cStr)
	C.myprint(cStr)

	var goStr string = C.GoString(cStr)
	fmt.Println("Go string: ", goStr)

	userStruct()
}

func userStruct() {
	var name *C.char = C.CString("Gopher")
	var age C.int = C.int(26)
	var user C.user = C.new_user(name, age)
	printUser(user)
	C.process_user(&user)
	printUser(user)
}

func printUser(user C.user) {
	fmt.Printf(`%T, user { 
	name (%p): '%s', 
	age: '%d' 
}
`, user, user.name, C.GoString(user.name), int(user.age))
}
