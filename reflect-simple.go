package main

import (
	"fmt"
	"reflect"
)

func main() {
	var (
		f float32 = 2.1
		i int16 = 3
		s []string
	)

	s = append(s, "hello")
	ref(0, 1.1, f, i, "hi",s, true)
}

func ref(a ...interface{}) {
	for argIndex, arg := range a {
		fmt.Println("[reflect] argNum: ", argIndex, ", arg:", reflect.ValueOf(arg), ", type:", reflect.TypeOf(arg))

		switch f := arg.(type) {
		case bool:
			fmt.Println("[reflect] type bool")
		case int:
			fmt.Println("[reflect] type int")
		case uintptr:
			fmt.Println("[reflect] type uintptr")
		case string:
			fmt.Println("[reflect] type string")
		case float64:
			fmt.Println("[reflect] type float64")
		case reflect.Value:
			fmt.Println("[reflect] type reflect.Value")
		default:
			fmt.Println("[reflect] unknown type of:", f)
		}
	}
}
