package main

import (
	"fmt"
)

func deferWithoutReturnName() int {
	var (
		i int
	)

	defer func() {
		i++
		fmt.Println("1st:", i)
	}()

	defer func() {
		i++
		fmt.Println("2nd:", i)
	}()

	return i
}

func deferWithReturnName() (i int) {
	defer func() {
		i++
		fmt.Println("1st:", i)
	}()

	defer func() {
		i++
		fmt.Println("2nd:", i)
	}()

	return i // return 0 效果是一样的
}

func deferWithAddress() *int {
	var (
		i int
	)

	defer func() {
		i++
		fmt.Println("1st:", i)
	}()

	defer func() {
		i++
		fmt.Println("2nd:", i)
	}()

	return &i
}

/**
* 从结果分析看：
* 当返回值命名时，使用的是地址方式的引用
 */
func main() {
	fmt.Println("deferWithoutReturnName", deferWithoutReturnName())
	fmt.Println("deferWithReturnName", deferWithReturnName())
	fmt.Println("deferWithAddress", *deferWithAddress())
}
