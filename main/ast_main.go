package main

import (
	"context"
	"fmt"
	"sync"
)

type AA map[string]interface{}

func main() {
	fmt.Println("2222")
	//ctx, cancel := context.WithCancel(context.Background())
	//defer cancel()
	//ctx = context.WithValue(ctx, "111", 222)
	//var group sync.WaitGroup
	//group.Add(4)
	//go A1(ctx, &group)
	//go A2(ctx, &group)
	//go A3(ctx, &group)
	//go A4(ctx, &group)
	//time.Sleep(time.Second)
	//cancel()
	//time.Sleep(time.Second)
	//group.Wait()
	//fmt.Println("Over!")
}

func A1(ctx context.Context, group *sync.WaitGroup) {
	defer group.Done()
	<-ctx.Done()
	fmt.Println("A1")

}

func A2(ctx context.Context, group *sync.WaitGroup) {
	defer group.Done()
	<-ctx.Done()
	fmt.Println("A2")
}

func A3(ctx context.Context, group *sync.WaitGroup) {
	defer group.Done()
	<-ctx.Done()
	fmt.Println("A3")
}

func A4(ctx context.Context, group *sync.WaitGroup) {
	defer group.Done()
	<-ctx.Done()
	fmt.Println("A4")
}
