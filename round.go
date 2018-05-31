package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	round()
}

func round() {
	startAt := 0
	open := 20
	close := 5

	for i := 0; i < 100; i = i + 2 {
		r := (i - startAt) % (open + close)
		fmt.Println(i, r < open)
	}
}
