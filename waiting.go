package main

import (
	"fmt"
	"time"
)

func waiter(t time.Duration) {
	for {
		for _, r := range "-`'‘-、," {
			fmt.Printf("hello  \r%c", r)
			time.Sleep(t)
		}
	}
}

func fib(n int) int {
	if n < 2 {
		return n
	}

	return fib(n-1) + fib(n-2)
}

func main() {
	go waiter(100 * time.Millisecond)

	const n = 41
	result := fib(n)

	fmt.Printf("result of %d is %d \n", n, result)
}
