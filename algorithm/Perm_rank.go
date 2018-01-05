package main

import (
	"fmt"
)

var list = []int{1,3,7,5,9,4,2,6,8}

func main() {
	perm(list, 0, 7)
}

func perm(l []int, k, m int) {
	if k == m {
		for i := 0; i <= m; i++ {
			fmt.Print(list[i])
		}
		fmt.Println("")
	} else {
		for i := k; i <= m; i++ {
			swap(k, i)
			perm(list, k + 1, m)
			swap(k, i)
		}
	}
}

func swap(a, b int) {
	list[a], list[b] = list[b], list[a]
}