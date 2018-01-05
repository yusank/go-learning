package main

import (
	//"fmt"
	"io/ioutil"
	"os/exec"
	//"time"
	"fmt"
)

func genarate() {
	cmd := exec.Command("/bin/sh", "-c", "crunch 1 6 ABCDEFGabcdefg0123456789")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}

	if err := cmd.Start(); err != nil {
		panic(err)
	}

	if err := cmd.Wait(); err != nil {
		fmt.Println("wait:", err.Error())
		return
	}
}

func main() {
	go genarate()
}