package main

import (
	"fmt"
	"reflect"
)

func ss(v reflect.Value, s string) error {
	switch v.Kind() {
	case reflect.Map:
		fmt.Printf("Parsing '%s' into: %v\n", s, v.Type())
		fmt.Println("Settable: ", v.CanSet())
		//fmt.Println(reflect.MakeMap(v.Type()))
		//v.Set(reflect.MakeMap(v.Type()))
		//for !endList(lex) {
		//	lex.consume('(')
		//	key := reflect.New(v.Type().Key()).Elem()
		//	read(lex, key)
		//	value := reflect.New(v.Type().Elem()).Elem()
		//	read(lex, value)
		//	v.SetMapIndex(key, value)
		//	lex.consume(')')
		//}
	}
	return nil
}
