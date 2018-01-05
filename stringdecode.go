package main

import (
	"fmt"
	"net/url"
	"strconv"
)

func main() {
	s := "%7B%22code%22%3A%22"
	qq, err := url.QueryUnescape(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(qq)
	u, err := url.PathUnescape(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(u)
	a := fmt.Sprint(`"`, u, `"`)
	q, err := strconv.Unquote(a)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%v", q)
	}
}
