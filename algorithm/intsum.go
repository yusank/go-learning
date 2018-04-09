package main

import (
	"fmt"
)

func main() {
	nums := []int{3, 2, 4, 8, 9}
	t := 12

	r := twoSum(nums, t)
	fmt.Println(r)
}

func twoSum(nums []int, target int) []int {
	var ns []int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				ns = append(ns, i, j)
				return ns
			}
		}
	}

	return nil
}
