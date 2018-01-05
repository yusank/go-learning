package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	num  [12]int
	s, b []int
	n    int
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("请输入一个1-12自然数：")
	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Print("there were error reading,exiting program")
		return
	}
	fmt.Sscan(input, &n)
	if n <= 0 || n > 12 {
		fmt.Println("错误！超出输入范围")
		return
	}

	for i := 0; i < 12; i++ {
		num[i] = i + 1
	}
	s := num[:n]
	fullcombine(s, n)
}

// fullcombine 求数列的全组合
func fullcombine(s []int, n int) {
	for M := 1; M <= n; M++ {
		m := M
		b := make([]int, M)
		combine(s, n, m, b, M)
	}
	fmt.Println("全组合的个数为：", fac(n))
}

// fac 求全组合的个数
func fac(n int) int {
	num := 2
	for i := 1; i < n; i++ {
		num *= 2
	}
	return num - 1
}

// combine 求数列1-n中任选m个元素的所有组合。
// s[1..n]表示候选集，n为候选集大小，n>=m>0。
// b[1..M]用来存储当前组合中的元素(这里存储的是元素下标)，
// M表示满足条件的一个组合中元素的个数，M=m，这两个参数仅用来输出结果。
func combine(s []int, n int, m int, b []int, M int) {
	for i := n; i >= m; i-- {
		b[m-1] = i - 1
		if m > 1 {
			combine(s, i-1, m-1, b, M)
		} else {
			for j := M - 1; j >= 0; j-- {
				fmt.Print(s[b[j]], " ")
			}
			fmt.Println("")
		}
	}
}
