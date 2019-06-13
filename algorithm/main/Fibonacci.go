package main

import (
	"fmt"
	"sync"
	"time"
)

var fMap map[int]int

// 存在递归消耗和大量重复计算
// 使得消耗少量空间后，for循环效率提升很高
func main() {
	Aa(40)
}

func Aa(i int) {
	fMap = make(map[int]int)
	t := time.Now()
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		fmt.Println("F value:		", F(i))
		fmt.Println("F:", time.Since(t))
	}()
	go func() {
		defer wg.Done()
		fmt.Println("FM value:		", FM(i))
		fmt.Println("FM:", time.Since(t))
	}()
	wg.Wait()
	fmt.Println("Over!")
}

func F(i int) int {
	if i < 2 {
		return 1
	}
	return F(i-1) + F(i-2)
}

func FM(i int) int {
	if i < 2 {
		return 1
	}
	fMap[0] = 1
	fMap[1] = 1
	for j := 2; j <= i; j++ {
		fMap[j] = fMap[j-1] + fMap[j-2]
	}
	return fMap[i]
}
