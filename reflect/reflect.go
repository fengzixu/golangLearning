package main

import (
	"fmt"
	"reflect"
	"strconv"
)

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
			display(fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name), v.Field(i))
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			display(fmt.Sprintf("%s[%s]", path, formatAtom(key)), v.MapIndex(key))
		}
	case reflect.Ptr:
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
			display(path+".Value", v.Elem())
		}
	default:
		fmt.Printf("%s = %s\n", path, formatAtom(v))
	}
}

func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Int, reflect.Int64, reflect.Int8, reflect.Int32, reflect.Int16:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.Chan, reflect.Func:
		return v.Type().String() + "0x" + strconv.FormatUint(uint64(v.Pointer()), 16)
	default:
		return v.Type().String() + " value"

	}
}

type A struct {
	Domain     string `json:"do"`
	StateSlice []string
	LineMap    map[string]int
	Ptr        *int
	Inter      interface{}
}

type F interface {
}

func main() {
	m := make(map[string]int)
	m["fuck"] = 2
	m["you"] = 1
	p := 3
	i := 4
	a := A{
		Domain:     "xuran.qbox.net",
		StateSlice: []string{"a", "b", "c", "d"},
		LineMap:    m,
		Ptr:        &p,
		Inter:      i,
	}

	display("a", reflect.ValueOf(a))
	var haha *int
	//haha = nil
	fmt.Println(reflect.ValueOf(haha).Elem().Kind())
	fmt.Println(reflect.ValueOf(haha))
	var hehe interface{}
	fmt.Println(reflect.ValueOf(hehe).IsValid())
	hehe = 3
	fmt.Println(reflect.ValueOf(hehe).Kind())
	fmt.Println(reflect.ValueOf(a).Field(4).Kind())
	fmt.Println(reflect.ValueOf(a).Field(4).Elem().Kind())

	fmt.Println(reflect.ValueOf(a).FieldByName("do"))
	return
}
