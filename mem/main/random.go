package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		for {
			select {
			case ch <- 0:
			case ch <- 1:
			}
		}
	}()
	i := 0
	for v := range ch {
		fmt.Print(v)
		i++
		if i > 20 {
			break
		}
	}
}
