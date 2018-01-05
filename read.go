package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	_, err := os.Open("./area.txt")

	if err != nil {
		fmt.Print(err)
	} else {
		_, er := ioutil.ReadFile("./area.txt")
		if er != nil {
			fmt.Print(er)
		}
		fmt.Print("readfile ok!")
	}
	fmt.Print("openfile ok!")
}
