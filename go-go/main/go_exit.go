package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

//通过传递信息提前让goroutine结束
func main() {
	wg := new(sync.WaitGroup)
	toExit := make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		isRun := true
		for {
			select {
			case i := <-toExit:
				fmt.Println("go exit: ", i)
				runtime.Goexit()
			default:
				if isRun {
					go func() {
						for {
							time.Sleep(1 * time.Second)
							fmt.Println("second num:		", time.Now().Second())
						}
					}()
					isRun = false
				}

			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			time.Sleep(1 * time.Second)
			fmt.Println("for:		", i)
			if i == 3 {
				toExit <- i
			}
		}
	}()
	wg.Wait()
	fmt.Println("game over")
}
