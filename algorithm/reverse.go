package main

import (
	"math"
	"fmt"
)

func main()  {
	testSlice := []int{123,456,-987}
	for _,v := range testSlice {
		fmt.Println(foo(v))
	}
}

func foo(a int) int {
	abs := math.Abs(float64(a))
	b := int(abs)
	if a == 0 || b > 2147483648 {
		return 0
	}

	// k := 0
	result := 0
	for {
		
		if b == 0 {
			break
		}
		result *= 10
		result += b%10

		b /= 10
	}

	if a < 0 {
		result = 0-result
	}
	return result
}