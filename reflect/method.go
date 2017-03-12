package main

import (
	"fmt"
	"reflect"
	"strings"
)

type T struct {
	Text string
}

func (t T) GetText() string {
	fmt.Println("get text")
	return t.Text
}

func PrintMethod(v reflect.Value) {
	if !v.IsValid() {
		fmt.Println("input params is invalid")
		return
	}

	for i := 0; i < v.NumMethod(); i++ {
		fmt.Printf("func (%s) %s %s\n", v.Type().String(),
			v.Type().Method(i).Name, strings.TrimPrefix(v.Method(i).Type().String(), "func"))
		fmt.Println(v.Method(i).Call([]reflect.Value{}))
	}

	return
}

func main() {
	a := T{Text: "haha"}
	PrintMethod(reflect.ValueOf(a))
	return
}
