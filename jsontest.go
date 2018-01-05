package main

import (
	"encoding/json"
	"fmt"
)

type Recive struct {
	Name  string
	Age   int
	Phone string
}

type Mobile struct {
	Tel []int
}

func user() {
	var info Recive
	m := []int{2, 2, 3, 3, 4, 5, 6}
	info.Name = "aabb"
	info.Age = 22
	s, err := generateJson(m)
	if err != nil {
		fmt.Println("got an error", err)
	}
	info.Phone = s

	fmt.Println(info)
}

func generateJson(a []int) (string, error) {
	var (
		s  Mobile
		ss Mobile
	)

	for i := 0; i < len(a); i++ {
		s.Tel = append(s.Tel, a[i])
	}
	b, err := json.Marshal(s)
	if err != nil {
		return "", err
	}

	err = json.Unmarshal(b, &ss)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(ss.Tel)
	return string(b), nil
}

func main() {
	user()
}
