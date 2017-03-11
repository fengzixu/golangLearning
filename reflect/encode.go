package main

import (
	"bytes"
	"fmt"
	"reflect"
)

func Encode(value reflect.Value, buff *bytes.Buffer) (err error) {
	switch value.Kind() {
	case reflect.Invalid:
		_, err = buff.Write([]byte("nil"))
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fmt.Fprintf(buff, "%d", value.Int())
	case reflect.String:
		fmt.Fprintf(buff, "%q", value.String())
	case reflect.Ptr:
		err = Encode(value.Elem(), buff)
	case reflect.Array, reflect.Slice:
		_, err = buff.WriteString("(")
		if err != nil {
			return
		}

		for i := 0; i < value.Len(); i++ {
			err = Encode(value.Index(i), buff)
			if err != nil {
				return
			}

			if i != value.Len()-1 {
				_, err = buff.WriteString(" ")
				if err != nil {
					return
				}
			}
		}

		_, err = buff.WriteString(")")
	case reflect.Struct:
		_, err = buff.WriteString("(")
		if err != nil {
			return
		}

		for j := 0; j < value.NumField(); j++ {
			if j > 0 {
				_, err = buff.WriteString(" ")
				if err != nil {
					return
				}
			}

			_, err = buff.WriteString("(")
			if err != nil {
				return
			}

			_, err = buff.WriteString(value.Type().Field(j).Name)
			if err != nil {
				return
			}

			_, err = buff.WriteString(" ")
			if err != nil {
				return

			}

			err = Encode(value.Field(j), buff)
			if err != nil {
				return
			}

			_, err = buff.WriteString(")")
			if err != nil {
				return
			}
		}

		_, err = buff.WriteString(")")
		if err != nil {
			return
		}

	case reflect.Map:
		_, err = buff.WriteString("(")
		if err != nil {
			return
		}

		for j, k := range value.MapKeys() {
			if j > 0 {
				_, err = buff.WriteString(" ")
				if err != nil {
					return
				}
			}

			_, err = buff.WriteString("(")
			if err != nil {
				return
			}

			err = Encode(k, buff)
			if err != nil {
				return
			}

			_, err = buff.WriteString(" ")
			if err != nil {
				return

			}

			err = Encode(value.MapIndex(k), buff)
			if err != nil {
				return
			}

			_, err = buff.WriteString(")")
			if err != nil {
				return
			}
		}

		_, err = buff.WriteString(")")
		if err != nil {
			return
		}

	default:
		fmt.Fprintf(buff, "%s:%s", "unsupport type", value.Kind())
	}

	return
}

type B struct {
	Fuck string `json:"fuck"`
	You  int    `json:"you"`
}

type A struct {
	Domain   string            `json:"domain"`
	Cache    []string          `json:"cache"`
	CdnState map[string]string `json:"cdnState"`
	Time     int               `json:"time"`
	FF       interface{}       `json:"ff"`
	B        `json:,inline`
}

func main() {
	b := B{
		Fuck: "fff",
		You:  3,
	}

	obj := A{
		Domain: "hahah.com",
		Cache:  []string{"1", "2", "3"},
		CdnState: map[string]string{
			"ws": "success",
			"tc": "success",
		},
		Time: 2,
		B:    b,
	}

	buff := bytes.Buffer{}
	err := Encode(reflect.ValueOf(obj), &buff)
	if err != nil {
		fmt.Printf("Error: %+v", err)
		return
	}

	fmt.Printf("%+v\n", string(buff.Bytes()))
	return
}
