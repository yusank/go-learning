package main

import (
	"fmt"
	"time"
)

const tt = "2017-09-01T12:23:34+08:00"
const ti = "2017-09-25T12:23:34+08:00"

func main() {
	fmt.Println(tt)
	a, b := time.Parse("2006-01-02T15:04:05Z07:00", tt)
	fmt.Println(a, b, a.Unix())
	c, _ := time.Parse("2006-01-02T15:04:05Z07:00", ti)
	fmt.Println(c.Unix())
	fmt.Println(time.Now().UnixNano()/int64(time.Millisecond))
	fmt.Println(time.Now().UnixNano()/int64(time.Microsecond))
	nowTime := time.Now()
	threeDaysAgo := time.Now().Add(4 * 24 * time.Hour)
	nn := 1506313414
	fromunix := time.Unix(int64(nn), 0)
	fmt.Println(">>>>",fromunix.String())
	fmt.Println(fromunix.Format("2006-01-02 15:04:05"))

	subs := threeDaysAgo.Sub(nowTime).Nanoseconds()

	fmt.Println(threeDaysAgo)
	h := time.Hour.Nanoseconds() * 24

	fmt.Println(h, ",,", subs)

	days := subs/h

	fmt.Println(days)

	//fmt.Println(threeDaysAgo, nowTime)

	//s := threeDaysAgo-nowTime

	//fmt.Println(s)
	//result := s / 3600
	//fmt.Println(result)
}
