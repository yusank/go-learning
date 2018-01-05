package main

import (
	"fmt"
	"strings"
//	"strconv"
	"time"
	//"unicode"
)

const(
	 abc = "2017-06-12T12:12:12+08:00"
	 aaa = "2017-06-11T12:13:14+08:00"
)
func main() {
	/*slice := strings.Split(abc, "-")
	j := strings.Join(slice, "")
	i, err := strconv.Atoi(j)
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Print(i)
		}
	j := strings.Join(slice[:2], "T")
	fmt.Println(j)
*/
	T()
}

func T() {
	t := time.Now().UTC()
	fmt.Print(t,"---------")
	tt , err := time.Parse("2006-01-02T15:04:05Z07:00", abc)
	fmt.Print(tt, "=============")
	if err != nil {
		fmt.Print(err)
	}
	tn, err := time.Parse("2006-01-02T15:04:05Z07:00", aaa)
	//tt := t.Format("2006-01-02")
	//d := t.Add(time.Duration(2) * time.Hour)
	dt := tn.Sub(t)
	fmt.Print(dt)
	dd := tt.Sub(t)
	fmt.Print(dd)
	dm := dt + dd
	fmt.Print(dm)
	//t, err := time.Parse("2006-01-02", "2017-06-11")
	//if err != nil {
	//	fmt.Print(err)
	//}
	//
	//fmt.Print(t)
	if strings.Contains("hello", "") {
		fmt.Println("success")
		}

}
