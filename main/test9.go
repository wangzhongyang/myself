package main

import (
	"fmt"
	"time"
)

type People2 struct {
	Name string
	Age  int
}

func main() {
	defer func() {
		for i := 0; i < 20; i++ {
			fmt.Println(time.Now().Format("2005-01-02 15:04:05"))
		}
	}()
	time.Sleep(100 * time.Second)
	fmt.Println("game over")
}
