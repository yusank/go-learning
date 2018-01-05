package main

import (
	"fmt"
	"regexp"
)

func main() {
	reg, err := regexp.Compile(`\Wid=(.*)`)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(reg.FindAllString("http://music.163.com/#/mv?id=1234", -1))
}
