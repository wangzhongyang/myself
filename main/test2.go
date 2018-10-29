package main

import (
	"fmt"
	"time"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()
		getPanicRecover()
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("-------1------")
	go getPanic()
	time.Sleep(1 * time.Second)
	fmt.Println("-------2------")
}

func getPanicRecover() {
	panic("this is panic recover")
}

func getPanic() {
	panic("this is panic")
}
