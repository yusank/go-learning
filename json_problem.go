package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	// get request
	r, _ := http.Get("https://api.robinhood.com/quotes/AAPL/")
	defer r.Body.Close()

	// get the []byte
	b, _ := ioutil.ReadAll(r.Body)

	// evidence the json was captured
	fmt.Printf("%s\n", b)

	// unmarshal to quote
	var quote Quote
	err := json.Unmarshal(b, &quote)
	if err != nil {
		fmt.Printf("got error %v \n", err)
	}

	fmt.Print("hello")
	// print results
	fmt.Printf("%s %s \n", quote.updated_at, quote.last_trade_price)
}

type Quote struct {
	updated_at, last_trade_price string
}
