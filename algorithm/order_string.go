package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "i am a good student."
	str := strspilt(s)
	result := order(str)
	fmt.Print(result)
}

func strspilt(s string) []string {
	return strings.Split(s, " ")
}

func order(s []string) string {
	m := len(s)

	for i := 0; i < m-1; i++ {
		t := s[i]
		s[i] = s[m-1]
		s[m-1] = t
		m--
	}

	return strings.Join(s, " ")
}
