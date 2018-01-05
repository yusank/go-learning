package main

import (
	"fmt"
	"reflect"
)

type Struct struct {
	a uint32
}

func main() {
	var (
		a uint32 = 5
		b        = &a
	)

	tryReflect(false, 5, 8.1, "hello", &Struct{}, Struct{}, a, &a, b)
}

func tryReflect(a ...interface{}) {
	var (
		s Struct
		c interface{} = (*int)(nil)
	)

	for argIndex, arg := range a {
		fmt.Println("[reflect] argNum: ", argIndex, ", arg:", reflect.ValueOf(arg), ", type:", reflect.TypeOf(arg), ",kind", reflect.TypeOf(arg).Kind())
	}

	fmt.Println(c == nil)
	fmt.Println(reflect.ValueOf(c).IsNil())

	fmt.Println("[reflect] Value of s:", reflect.ValueOf(s), ", type of s:", reflect.TypeOf(s))
	fmt.Println("[reflect]:", reflect.ValueOf(s).FieldByName("a").Type())
}
