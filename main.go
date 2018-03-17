package main

import "fmt"

func main() {
	if 100%10 == 0 {
		fmt.Println("yes")
	}
	a := "hello_"
	c := a[len(a)-1]
	if string(c) == "_" {
		fmt.Println("nice")
	}

	fmt.Println(a)
	foo(a)
	fmt.Println(a)

	d := 0
	bar(d)
	fmt.Println(d)
}

func foo(a string) {
	a = a + "_"
}

func bar(a int) {
	a++
}
