package main

import (
	"fmt"
)

func main() {
	nonBlocked()
	nonCacheChannelGeneral()
	blocked()
}

func nonBlocked() {
	ch := make(chan int, 1)

	ch <- 0

	select {
	case <-ch:
		fmt.Println("[nonBlocked]:read from channel")
	default:
		fmt.Println("[nonBlocked]:no data")
	}
}

func blocked() {
	ch := make(chan int)

	ch <- 0

	select {
	case <-ch:
		fmt.Println("[blocked]:read from channel")
	default:
		fmt.Println("[blocked]:no data")
	}
}

func nonCacheChannelGeneral() chan<- int {
	ch := make(chan int)

	go func() {
		select {
		case <-ch:
			fmt.Println("[nonCacheChannelGeneral]:read from channel")
		default:
			fmt.Println("[nonCacheChannelGeneral]:no data")
		}
	}()

	ch <- 0

	return ch
}
