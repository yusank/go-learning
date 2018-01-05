package main

import (
	"fmt"
	"time"
)

type Test struct {
	Timestamp time.Time
	Count     int
	Times     func()
}

func (this *Test) Timeing() {
	for i := 0; i < 10; i++ {
		fmt.Println(this.Timestamp.Unix())
		fmt.Println(this.Count)
	}
}

func (this *Test) Counting() {
	var a Test
	var t int

	t += 1
	a = Test{
		Count: 1,
	}
}
func init() {
	Test.Times = Test.Counting()
}
func main() {
	Test.Timeing()
}
