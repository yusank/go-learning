package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	var stat syscall.Statfs_t
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	syscall.Statfs(wd, &stat)
	fmt.Println(stat.Bavail * uint64(stat.Bsize))
}
