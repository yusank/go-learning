package main

import (
	"log"
	"net/http"
)



func main() {
	http.Handle("/",http.FileServer(http.Dir("/Users/yusank/workspace/JS/wxbackground-client")))

	err := http.ListenAndServe(":4000", nil)
	if err != nil {
		log.Fatal(err)
	}
}