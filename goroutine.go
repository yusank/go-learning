package main

import (
	"fmt"
	"time"
)

func generate(out chan<- int) {
	for x := 0; x < 100; x++ {

		out <- x
		fmt.Println("generate-->", x)
	}

	close(out)
}

func produce(out chan<- int, in <-chan int) {
	for x := range in {
		time.Sleep(time.Millisecond * 100)
		out <- x * x
		fmt.Println("produce-->", x)

	}
	close(out)
}

func printer(in <-chan int) {
	for v := range in {
		fmt.Println("print>--", v)
	}
}

func main() {
	num := make(chan int)
	res := make(chan int)

	go generate(num)
	go produce(res, num)

	printer(res)
}
