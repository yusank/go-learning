package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	fmt.Println("Is channel closed:", isChannelClosedWithRecover(ch))
	close(ch)
	fmt.Println("Is channel closed:", isChannelClosedWithRecover(ch))
}

func isChannelClosedWithRecover(ch chan int) (closed bool) {
	defer func() {
		// 如果从 panic 中恢复，p 为非 nil 值；
		if p := recover(); p != nil {
			fmt.Println("Panic recovered:", p)
			closed = true // 修正关闭状态
		}
	}()

	select {
	case ch <- 0: // 向关闭的 Channel 发送内容，会触发 panic
	default: // default 使 Channel 不阻塞在此

	}

	return closed
}
