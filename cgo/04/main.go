package main

import "fmt"

func main() {
	fmt.Println("===================== callPyFnNoArgs()")
	callPyFnNoArgs()
	fmt.Println("===================== callPyFnWithArgs()")
	callPyFnWithArgs("Mark & Paul")
	fmt.Println("===================== multipleCgoCalls()")
	multipleCgoCalls("Luke")
}
