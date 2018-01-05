package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"time"
	"strings"
)

func main() {
	begin := time.Now()
	client := &http.Client{}

	req, err := http.NewRequest("POST", "http://127.0.0.1:3010/publish", strings.NewReader("gameId=test"))
	if err != nil {
		// handle error
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)


	defer func() {
		end := time.Now()
		fmt.Println(end.Sub(begin))
		resp.Body.Close()
	}()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("got error:", err)
	}

	fmt.Println(string(body))
}
