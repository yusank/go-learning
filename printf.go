package main

import "fmt"

func main() {
	sl := []string{"a","c","f"}
	fmt.Printf("hello(%s)",sl)
	aa := fmt.Sprintf("a %s", "b")
	fmt.Print(aa)
	fmt.Println()
	fmt.Printf("%t", true)
}
