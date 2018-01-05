package main

import (
	"fmt"
)

func main() {
	fmt.Print(calculate(3, 2))
}

func calculate(n, m int) int {
	if n == 0 {
		return 1
	}

	if n >= 1 {
		if m == 0 {
			if n == 1 {
				return 2
			} else {
				return n + 2
			}
		} else {
			return calculate(calculate(n - 1, m), m - 1)
		}
	}

	return -1
}