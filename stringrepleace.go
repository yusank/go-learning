package main

import (
	"fmt"
	"reflect"
	"time"
)

type NotknownType struct {
	s1 [2]time.Time
	s2 int
	s3 uint64
}

func main() {
	secret := NotknownType{
		s2: 11,
		s3: 22,
	}
	value := reflect.ValueOf(secret)
	for i := 0; i < value.NumField(); i++ {
		//bv := value.Field(i)
		s := fmt.Sprint(value.Field(i).Type())
		switch s {
		case "[2]time.Time":

			fmt.Println("aa")

			fmt.Println(value.Field(i).Type())
			fmt.Println("it`s a time")
		case "uint64", "int":
			fmt.Println(value.Field(i).Type())
		}
	}
}
