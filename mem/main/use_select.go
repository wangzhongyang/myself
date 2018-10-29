package main

import (
	"fmt"
	"time"
)

func worker(cannel chan bool) {
	for {
		select {
		default:
			fmt.Println("hello")
			// 正常工作
		case <-cannel:
			fmt.Println("exit 0")
			// 退出
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	cannel := make(chan bool)
	go worker(cannel)

	time.Sleep(time.Second)
	cannel <- true
	time.Sleep(time.Second)
}
