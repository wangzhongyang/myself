package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

// 协程中有错误，main函数退出

func main() {
	var wg sync.WaitGroup
	errChan := make(chan error, 1)
	wg.Add(1)
	time3 := 3
	go CreateErr(errChan, &wg, time3, func(a int) error {
		time.Sleep(time.Duration(time3) * time.Second)
		fmt.Printf("go func %ds\n", time3)
		//return errors.New("this is 3s error")
		return nil
	})
	wg.Add(1)
	time2 := 2
	go CreateErr(errChan, &wg, time2, func(a int) error {
		time.Sleep(time.Duration(time2) * time.Second)
		fmt.Printf("go func %ds\n", time2)
		return errors.New("this is 2s error")
		//return nil
	})

	for {
		select {
		case err := <-errChan:
			fmt.Println("this is select case err:		", err.Error())
			return
		case <-time.After(4 * time.Second):
			fmt.Println("time out")
			return
		}
	}
	if err := <-errChan; err != nil {
		fmt.Println(err.Error())
		return
	}

	wg.Wait()
}

func CreateErr(errChan chan error, wg *sync.WaitGroup, a int, callback func(a int) error) {
	err := callback(a)
	defer func() {
		wg.Done()
		if err := recover(); err != nil {
			errChan <- errors.New(fmt.Sprintf("this is recover err:		%s", err.(string)))
		}
	}()
	if err != nil {
		errChan <- err
	}
	return
}
