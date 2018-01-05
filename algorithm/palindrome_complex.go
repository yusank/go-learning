package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	inputRead *bufio.Reader
	lastStr   string
	lastInt   int
)

func readInput() string {
	inputRead = bufio.NewReader(os.Stdin)
	fmt.Print("input an string...:")
	input, err := inputRead.ReadString('\n')
	if err != nil {
		return ""
	}

	return input[:len(input)-1]
}

func pali(s string) bool {
	n := len(s)

	if n < 1 {
		return false
	}

	if n == 1 {
		count(s)
	}

	for i := 0; i < n-1; i++ {
		for j := n - 1; j > i; j-- {
			if lastInt > len(s)/2 {
				return true
			}
			count(s[i : j+1])
		}
	}

	if lastInt == 0 {
		return false
	}

	return true
}

func count(str string) {
	n := len(str)

	if n == 1 {
		if n > lastInt {
			lastInt = n
			lastStr = str
		}
		return
	}

	for i := 0; i < n-1; i++ {
		if str[i] == str[n-1] {
			n--
			continue
		} else {
			return
		}
	}

	if len(str) > lastInt {
		lastInt = len(str)
		lastStr = str
	}

	return
}

func main() {
	ok := pali(readInput())
	if ok {
		fmt.Printf("last palindrome is %s and length is %d\n", lastStr, lastInt)
	} else {
		fmt.Print("no palindrome.")
	}
}
