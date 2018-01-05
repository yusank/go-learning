package main

import (
	"github.com/google/go-github/github"
	//"github.com/pkg/errors"
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/oauth2"
	//"net/http"
)

const (
	token           = "a303d2587290e6573e1f7b50ce0b11cc8b747dac"
	invalidTokenErr = "401 Unauthorized"
	//token = ""
)

var ErrWrongToken = errors.New("GET https://api.github.com/rate_limit: 401 Bad credentials []")

func newClient() {
	tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	//tokenSource := *new(oauth2.TokenSource)
	httpClient := oauth2.NewClient(oauth2.NoContext, tokenSource)
	client := github.NewClient(httpClient)

	req, er := client.NewRequest("GET", "", nil)
	if er != nil {
		fmt.Println("er error", er, req)
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		fmt.Println("error:", err)
		fmt.Println(resp)
	}
	if resp.Header.Get("Status") == invalidTokenErr {
		fmt.Println("all is done.", resp)
	} else {
		fmt.Println(resp.Header.Get("Status"))
		fmt.Println(resp)
	}
}

func main() {
	newClient()
	// clinet := newClient()
	// response := new(struct {
	// 	R *github.RateLimits `json:"resources"`
	// })
	// req, er := clinet.NewRequest("GET", "rate_limit", nil)
	// if er != nil {
	// 	fmt.Println(er)
	// }
	// resp, err :=
	// if err != nil {
	// 	fmt.Println("error:", resp, err)
	// } else {
	// 	fmt.Println("reps:", resp)
	// }
}
