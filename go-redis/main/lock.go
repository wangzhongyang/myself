package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/gomodule/redigo/redis"
)

var lock sync.Mutex

func main() {
	client, err := redis.Dial("tcp", ":34560")
	if err != nil {
		fmt.Println("conn error:	" + err.Error())
	}
	defer client.Close()

	client.Do("set", "test", "30")
	fmt.Println(redis.Int(client.Do("get", "test")))
	arr := []int{1, 2, 3}
	for _, v := range arr {
		go func() {
			defer func() {
				if err := recover(); err != nil {
					fmt.Println("error:		", err.(string))
				}
			}()
			fmt.Println("go:			", v)
			for j := 0; j < 5; j++ {
				lock.Lock()
				a, err := redis.Int(client.Do("get", "test"))
				if err != nil {
					fmt.Println("get redis error:		", err.Error())
					break
				}
				a = a + 1
				fmt.Println(v, j, a)
				_, err = client.Do("set", "test", a)
				if err != nil {
					fmt.Println("set redis error:		", err.Error())
					break
				}
				lock.Unlock()
			}
		}()
	}
	time.Sleep(5 * time.Second)
	c, _ := redis.Int(client.Do("get", "test"))
	fmt.Println("end:	", c)
}
