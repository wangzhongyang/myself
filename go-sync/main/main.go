package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

func main() {
	SyncOnce()
}

func A() {
	var wg sync.WaitGroup
	errorChan := make(chan error)
	go func() {
		wg.Add(1)
		time.Sleep(3 * time.Second)
		fmt.Println("sleep 3 s")
		errorChan <- errors.New("error sleep 3 s")
		wg.Done()
	}()

	go func() {
		wg.Add(1)
		time.Sleep(5 * time.Second)
		fmt.Println("sleep 5 s")
		errorChan <- errors.New("error sleep 5 s")
		wg.Done()
	}()

	fmt.Println("main 1")
	wg.Wait()
	if err := <-errorChan; err != nil {
		fmt.Println("error is :		", err.Error())
	}
	fmt.Println("main 2")
}
