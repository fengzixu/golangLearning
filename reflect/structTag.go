package main

import (
	"fmt"
	"github.com/qiniu/errors"
	"net/http"
	"reflect"
	"strconv"
)

func ExtractHttpReq(req *http.Request, ptr interface{}) error {
	m := make(map[string]reflect.Value)
	value := reflect.ValueOf(ptr).Elem()
	if value.Kind() != reflect.Struct || !value.IsValid() {
		return errors.New("input data is invalid")
	}

	for i := 0; i < value.NumField(); i++ {
		tag := value.Type().Field(i).Tag.Get("http")
		if tag == "" {
			tag = value.Type().Field(i).Name
		}
		m[tag] = value.Field(i)
	}

	if err := req.ParseForm(); err != nil {
		return err
	}

	for key, value := range req.Form {
		f, ok := m[key]
		if !ok || !f.IsValid() {
			continue
		}

		for _, item := range value {
			Populate(f, item)
		}
	}

	return nil
}

func Populate(value reflect.Value, item string) {
	switch value.Kind() {
	case reflect.Slice:
		vv := reflect.New(value.Type().Elem()).Elem()
		Populate(vv, item)
		value.Set(reflect.Append(value, vv))
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Uint64:
		num, _ := strconv.ParseInt(item, 10, 64)
		value.SetInt(num)
	case reflect.Bool:
		num, _ := strconv.ParseBool(item)
		value.SetBool(num)
	case reflect.String:
		value.SetString(item)
	default:
		fmt.Println("unsupport type")
	}

	return
}

func main() {
	fmt.Println("Func main....")
	return
}
