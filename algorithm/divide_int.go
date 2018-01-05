package main

import (
	"fmt"
)

func divide(n, m int) int {
	if n < 1 {
		return 0
	}

	if n == 1 || m == 1 {
		return 1
	}

	if n < m {
		return divide(n, n)
	}

	if n == m {
		return divide(n, m-1) + 1
	}

	return divide(n, m-1) + divide(n-m, m)
}

func main() {
	fmt.Print(divide(10, 5))
}
