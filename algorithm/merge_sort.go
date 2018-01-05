package main

import (
	"fmt"
)

func mergeSort(sli []int) []int {
	if len(sli) < 2 {
		return sli
	}

	i := len(sli) / 2
	left := sli[:i]
	right := sli[i:]

	return merge(mergeSort(left), mergeSort(right))
}

func merge(left, right []int) []int {
	var result []int
	for len(left) != 0 && len(right) != 0 {
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}

	for len(left) != 0 {
		result = append(result, left[0])
		left = left[1:]
	}

	for len(right) != 0 {
		result = append(result, right[0])
		right = right[1:]
	}

	return result
}

func main() {
	i := []int{2, 4, 6, 9, 8, 5, 1, 3}
	fmt.Printf("before sort %v \n", i)

	fmt.Printf("after sort %v \n", mergeSort(i))
}
