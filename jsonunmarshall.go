package main

import (
	"encoding/json"
	"fmt"
)

type A struct {
	CategoryID []uint8 `json`
}

func main() {

	var a A

	s := "{'CategoryID': [1, 2, 3, 4]}"

	b := []byte(s)

	err := json.Unmarshal(b, &a)

	if err != nil {
		fmt.Print(err)
	}

	fmt.Print(a)

}
