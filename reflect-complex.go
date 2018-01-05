package main

import(
	"fmt"
	"reflect"
)

type foo struct {
	X int
	Y int
}

func c() {
	f := foo{
		X: 1,
		Y: 2,
	}
	p := foo{
		X: 3,
		Y: 4,
	}

	fi := reflect.ValueOf(f).Field(0)
	fx := reflect.ValueOf(f).Field(1)

	pi := reflect.ValueOf(&p)
	// Elem returns the value that the interface v contains
	// or that the pointer v points to.
	// It panics if v's Kind is not Interface or Ptr.
	// It returns the zero Value if v is nil.
	px := reflect.ValueOf(&p).Elem().Field(1) 

	fmt.Println(fi.CanSet(), fi, fi.Type())
	fmt.Println(fx.CanSet(), fx, fx.Type())
	fmt.Println(pi.CanSet(), pi, pi.Type())
	fmt.Println(px.CanSet(), px, px.Type())

	px.SetInt(22)
	fmt.Println(reflect.ValueOf(p).Field(1))
}

func main() {
	c()
}
