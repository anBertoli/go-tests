package main

/*
#cgo CFLAGS: -g -Wall -I/usr/local/include -I/usr/local/Cellar/python@3.8/3.8.15/Frameworks/Python.framework/Versions/3.8/include/python3.8
#cgo LDFLAGS: -L/usr/local/lib -L/usr/local/opt/python@3.8/Frameworks/Python.framework/Versions/3.8/lib/python3.8/config-3.8-darwin -lpython3.8
#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include <ctype.h>
#include "./c/glue.c"
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func init() {
	C.init_python()
}

func callPyFnNoArgs() {
	fileName := C.CString("my_test")
	funcName := C.CString("get_num")
	defer func() {
		C.free(unsafe.Pointer(fileName))
		C.free(unsafe.Pointer(funcName))
	}()

	var num C.long = C.call_my_func_ret_num(fileName, funcName)
	if num == 0 {
		fmt.Println("error call_my_func_ret_num: ", C.GoString(C.py_last_error()))
		return
	}
	fmt.Printf("result, type: '%T', val: '%+v'\n", num, num)
}

func callPyFnWithArgs(name string) {
	fileName := C.CString("my_test")
	funcName := C.CString("concat")
	funcArg := C.CString(name)
	defer func() {
		C.free(unsafe.Pointer(fileName))
		C.free(unsafe.Pointer(funcName))
		C.free(unsafe.Pointer(funcArg))
	}()

	var res *C.char = C.call_my_func_ret_str(fileName, funcName, funcArg)
	if res == nil {
		fmt.Println("error: ", C.GoString(C.py_last_error()))
		return
	}
	fmt.Printf(
		"result, type: '%T', val: '%+v'\n",
		res, C.GoString(res),
	)
}

func multipleCgoCalls(name string) {
	fileName := C.CString("my_test")
	funcName := C.CString("concat")
	funcArg := C.CString(name)
	defer func() {
		C.free(unsafe.Pointer(fileName))
		C.free(unsafe.Pointer(funcName))
		C.free(unsafe.Pointer(funcArg))
	}()

	var module *C.PyObject = C.PyImport_ImportModule(fileName)
	if module == nil {
		fmt.Println("error PyImport_ImportModule: ", C.GoString(C.py_last_error()))
		return
	}

	var fn *C.PyObject = C.PyObject_GetAttrString(module, funcName)
	if fn == nil {
		fmt.Println("error PyObject_GetAttrString: ", C.GoString(C.py_last_error()))
		return
	}

	// build function args and call function
	var args *C.PyObject = C.PyTuple_New(1)
	var ss *C.PyObject = C.PyUnicode_FromStringAndSize(
		funcArg,
		C.long(C.strlen(funcArg)),
	)
	C.PyTuple_SetItem(args, 0, ss)

	var out *C.PyObject = C.PyObject_CallObject(fn, args)
	if out == nil {
		return
	}

	// runtime check
	var x C.int = C.PyUnicode_Check(out)
	fmt.Println(x)
	//switch {
	//case C.PyUnicode_Check(out) == C.int(1):
	//	fmt.Println("is a a string")
	//default:
	//	fmt.Println("unhandled")
	//}

	var some *C.PyObject = C.PyUnicode_AsUTF8String(out)
	if some == nil {
		fmt.Println("PyUnicode_AsUTF8String error\n")
		return
	}

	var str *C.char = C.PyBytes_AsString(some)
	if str == nil {
		fmt.Println("PyBytes_AsString error\n")
		return
	}

	goStr := C.GoString(str)
	fmt.Printf(
		"result, type: '%T', val: '%+v'\n",
		str, goStr,
	)
}
