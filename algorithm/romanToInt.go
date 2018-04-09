package main

import (
	"strings"
	"fmt"
)

var roman = map[string]int{
	"i": 1,
	"v": 5,
	"x": 10,
	"l": 50,
	"c": 100,
	"d": 500,
	"m": 1000,
}

func main() {
	fmt.Println(romanToInt("XLV"))
}

func romanToInt(r string) int {
	var result int
	if r == "" {
		return 0
	}

	r = strings.ToLower(r)
	bytes := []byte(r)
	for i,b := range bytes { 
		if  i == 0 || (roman[string(b)] <= roman[string(bytes[i-1])]){
			result += roman[string(b)]
		} else {
			result += roman[string(b)] - roman[string(bytes[i-1])]*2
		}
	}

	return result
}
