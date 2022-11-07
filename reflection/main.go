package main

import (
	"fmt"
	"reflect"
	"strconv"
)

// Go provides a mechanism to update variables and inspect their values at run time,
// to call their methods, and to apply the operations intrinsic to their representation,
// all without knowing their types at compile time. This mechanism is called reflection.
// Reflection also lets us treat types themselves as first-class values.

func main() {
	//fmt.Println("============= NUMS EXAMPLE =============")
	//exNums()
	//fmt.Println("\n\n")
	//
	//fmt.Println("============= FUNC EXAMPLE =============")
	//exFunc()
	//fmt.Println("\n\n")
	//
	//fmt.Println("============= STRUCT EXAMPLE =============")
	//exStruct()
	//fmt.Println("\n\n")

	x := map[string]int{}
	ss(reflect.ValueOf(x), "")
}

// //////////////////////////////////////////////////////////////////
// Reflection on numbers
func exNums() {
	var t1 reflect.Type = reflect.TypeOf(3)

	fmt.Printf("t: '%v', ", t1)
	fmt.Printf("t.String(): '%v', ", t1.String())
	fmt.Printf("t.Kind(): '%v', ", t1.Kind())
	fmt.Printf("t.Name(): '%v', ", t1.Name())
	fmt.Printf("t.Align(): '%v', ", t1.Align())
	fmt.Printf("t.Size(): '%v', ", t1.Size())
	fmt.Printf("t.Bits(): '%v', ", t1.Bits())
	fmt.Printf("t.Kind(): '%v', ", t1.Kind())
	fmt.Printf("t.PkgPath(): '%v'\n", t1.PkgPath())
}

// //////////////////////////////////////////////////////////////////
// Reflection on functions

func myFunc(a []string, b uint) (*int, error) { return nil, nil }

func exFunc() {
	var t1 reflect.Type = reflect.TypeOf(myFunc)
	inspectFunc(0, t1)
}

func inspectFunc(pad uint, t reflect.Type) {
	inspectType(pad, t)

	fmt.Printf("\n=== Function params\n")
	fmt.Printf("t.NumIn(): '%v'\n", t.NumIn())
	for i := 0; i < t.NumIn(); i++ {
		var param reflect.Type = t.In(i)
		fmt.Printf("t.ParamIn(%d): '%+v'\n", i, param.String())
		inspectType(2, param)
	}

	fmt.Printf("\n=== Function outs\n")
	fmt.Printf("t.NumOut(): '%v'\n", t.NumOut())
	for i := 0; i < t.NumOut(); i++ {
		var out reflect.Type = t.Out(i)
		fmt.Printf("t.ParamOut(%d): '%+v'\n", i, out.String())
		inspectType(2, out)
	}
}

// //////////////////////////////////////////////////////////////////
// Reflection on structs

type someStruct struct {
	A uint32 `test:"test_val"`
	a int
	b string
	float64
}

func (ss someStruct) unexported(a int, b string) []int    { return nil }
func (ss someStruct) Exported(a float64, b []string) *int { return nil }

func exStruct() {
	var t reflect.Type = reflect.TypeOf(someStruct{4, 2, "ehy", 3})
	inspectStruct(0, t)
}

func inspectStruct(pad uint, t reflect.Type) {
	inspectType(pad, t)
	fmt.Printf("t.FieldAlign(): '%v', ", t.FieldAlign())
	fmt.Printf("t.NumField(): '%v'\n", t.NumField())

	// Inspect struct fields.
	fmt.Println("\n=== Struct fields")
	for i := 0; i < t.NumField(); i++ {
		var f reflect.StructField = t.Field(i)
		fmt.Printf("t.Field(%d): '%+v'\n", i, f)
		fmt.Printf("  exported: %v\n", f.IsExported())
		val, found := f.Tag.Lookup("test")
		if found {
			fmt.Printf("  'test' tag found: '%v'\n", val)
		}
		inspectType(2, f.Type) // recursive
	}

	// Inspect struct methods.
	fmt.Println("\n=== Struct methods")
	fmt.Printf("t.NumMethod(): '%v'\n", t.NumMethod())
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("t.Method(%d): '%+v' -> exported: %v\n", i, m, m.IsExported())
		inspectFunc(2, m.Type)
	}
}

//func ex1() {
//
//	// Because reflect.TypeOf returns an interface value’s dynamic type,
//	// it always returns a concrete type.
//	var t1 reflect.Type = reflect.TypeOf(3) // a reflect.Type
//
//	fmt.Println(t1.String()) // "int"
//	fmt.Println(t1)          // "int"
//
//	var t2 reflect.Type = reflect.TypeOf(someStruct{}) // a reflect.Type
//	fmt.Println(t2.String())                           // "main.someStruct"
//	fmt.Println(t2)
//}

func display(path string, v reflect.Value) {
	switch v.Kind() {
	case reflect.Invalid:
		fmt.Printf("%s = invalid\n", path)
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			display(fmt.Sprintf("%s[%d]", path, i), v.Index(i))
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			fieldPath := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)
			display(fieldPath, v.Field(i))
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			display(
				fmt.Sprintf("%s[%s]", path, formatAtom(key)),
				v.MapIndex(key),
			)
		}
	case reflect.Pointer:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			display(fmt.Sprintf("(*%s)", path), v.Elem())
		}
	case reflect.Interface:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			fmt.Printf("%s.type = %s\n", path, v.Elem().Type())
			display(path+".value", v.Elem())
		}
	default:
		// basic types, channels, funcs
		fmt.Printf("%s = %s\n", path, formatAtom(v))
	}
}

func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
	// ...floating-point and complex cases omitted for brevity...
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Slice, reflect.Map:
		return v.Type().String() + " 0x" +
			strconv.FormatUint(uint64(v.Pointer()), 16)
	default:
		// reflect.Array, reflect.Struct, reflect.Interface
		return v.Type().String() + " value"
	}
}

func inspectType(pad uint, t reflect.Type) {
	padS := padSpace(pad)
	fmt.Printf("%sInspecting type: '%v' --> ", padS, t)
	fmt.Printf("%st.Kind(): '%v', ", padS, t.Kind())
	fmt.Printf("%st.Name(): '%v', ", padS, t.Name())
	fmt.Printf("%st.Align(): '%v', ", padS, t.Align())
	fmt.Printf("%st.Size(): '%v', ", padS, t.Size())
	fmt.Printf("%st.Kind(): '%v', ", padS, t.Kind())
	fmt.Printf("%st.PkgPath(): '%v'\n", padS, t.PkgPath())
}

func padSpace(pad uint) string {
	s := ""
	for i := uint(0); i < pad; i++ {
		s += " "
	}
	return s
}
