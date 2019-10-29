package main

import (
	"fmt"
	"sync"
	"time"
)

var once sync.Once

// once中代码只执行一次，其他调用者等待完成后执行其他逻辑
func SyncOnce() {
	for i := 0; i < 5; i++ {
		go SyncOnceBody(i)
	}
	time.Sleep(3 * time.Second)
}

func SyncOnceBody(i int) {
	fmt.Println("coming body,i is:", i, getNow())
	once.Do(func() {
		fmt.Println("i am in doing once,sleep 1s,i is:", i, getNow())
		time.Sleep(time.Second)
	})
	fmt.Println("out body,i is:", i, getNow())
}

func getNow() string {
	return time.Now().Format("2006-01-02 15:04:05.000")
}
