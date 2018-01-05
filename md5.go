package main

import (
	"crypto/md5"
	"fmt"
	"os"
)

func main() {
	hash := md5.New()
	hash.Write([]byte(os.Args[1]))
	fmt.Printf("%x", hash.Sum(nil))
}
