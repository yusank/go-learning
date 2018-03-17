package main

import (
	"fmt"
	"time"
)

const (
	format1   = "0601"
	longForm  = "Jan 2, 2006 at 3:04pm (MST)"
	shortForm = "2006-Jan-02"
	newFormat = "2006-01-02"
)

func exampleDate() {
	t := time.Now()
	fmt.Printf("now is : %s\n", t)
	fmt.Println(t.AddDate(0, 4, 0))
}

// 把时间格式转换成给定的格式
func exampleTimeFormat() {
	fmt.Println("time is ", time.Now().Format(format1))
	fmt.Println("time is ", time.Now().Format(longForm))
	fmt.Println("time is ", time.Now().Format(shortForm))
}

// 传入的字符串解析并转换成默认的 2017-06-02 00:00:00 +0000 UTC 格式
// 但是传入格式与 parse 的第一参数的格式得一致。
func exmapleParse() {
	t, _ := time.Parse(format1, "2017-Jun-02")
	fmt.Println(t)
	newtime := t.AddDate(0, 0, 40)
	_, mon, _ := newtime.Date()
	fmt.Println(mon)

	n, err := time.Parse(newFormat, "2017-09-01")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(">>", n)
}

func NextMonthStr() string {
	t := time.Now()
	year, month, _ := t.Date()
	hour, min, sec := t.Clock()
	years := 0
	if month == time.December {
		years = 1
	}
	nt := time.Date(year+years, month+time.Month(1), 1, hour, min, sec, 0, t.Location())

	return nt.Format("200601")
}

func main() {
	exampleDate()
	exampleTimeFormat()
	exmapleParse()
	now := time.Now()
	start := time.Date(now.Year(), now.Month(), now.Day()-31, 0, 0, 0, 0, time.Local)
	fmt.Println(start)
	ttt := time.Now().AddDate(0, 1, 0)
	fmt.Println(ttt)
	fmt.Println(time.Now())
	fmt.Println(NextMonthStr())
}
