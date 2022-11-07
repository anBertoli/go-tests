package main

/*
#cgo CFLAGS: -g -Wall -I/usr/local/include
#cgo LDFLAGS: -L/usr/local/lib  -lss
#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include <ctype.h>
#include "./c/hello.c"
#include "ss.h"

void myprint(char* s) {
	printf("%s", s);
}

void print_ss(ss str) {
	printf("{len: %zu, free: %zu, buf (%p): %s}\n",
		str->len,
		str->free,
		str->buf,
		str->buf
	);
}
*/
import "C"

import (
	"fmt"
	"unsafe"
)

func main() {
	// Create new string using the ss library.
	// ss ss_new_from_raw(const char *init);
	rawStr := C.CString("My library.")
	var str C.ss = C.ss_new_from_raw(C.CString("My library."))
	defer C.free(unsafe.Pointer(rawStr))
	defer C.ss_free(str)
	if str == nil {
		fmt.Println("error")
		return
	}

	// Print the starting string.
	fmt.Printf("C.ss types in Go: {ss->len = %T, ss->free = %T, ss->buf = %T}\n",
		str.len, str.free, str.buf)
	C.print_ss(str)

	// void ss_to_upper(ss s);
	C.ss_to_upper(str)
	C.print_ss(str)

	// ss_err ss_concat_raw(ss s1, const char *s2);
	//rawConcat := C.CString("My library.")
	rawConcat := C.CString(" Consider a book.")
	defer C.free(unsafe.Pointer(rawConcat))

	var err C.ss_err = C.ss_concat_raw(str, rawConcat)
	if err != 0 {
		fmt.Println("error")
		return
	}
	fmt.Printf("%T, %+v\n", err, err)
	C.print_ss(str)
}
