package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	input       string
	inputReader *bufio.Reader
)

func readFromInput() string {
	inputReader = bufio.NewReader(os.Stdin)
	fmt.Print("input an string...:")
	input, err := inputReader.ReadString('\n')
	if err != nil {
		return ""
	}

	return input
}

func palindrome(s string) bool {
	str := s[:len(s)-1]
	n := len(str)

	if n < 1 {
		return false
	}

	if n == 1 {
		return true
	}

	for i := 0; i < n-1; i++ {
		if s[i] == s[n-1] {
			n--
			continue
		} else {
			return false
		}
	}

	return true
}

func main() {
	ok := palindrome(readFromInput())

	fmt.Print(ok)
}
