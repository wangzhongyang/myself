package go_groutine_pool

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

var chanIds chan string
var (
	A string = "www.baidu.com"
	B string = "www.google.com"
)

// TestPool
// 需求：每个URL未结束之前不可以重复执行
// 结果：输出三个 www.baidu.com，两个www.google.com
func TestPool(t *testing.T) {
	urls := []string{
		A,
		B,
	}
	chanIds = make(chan string, 10)
	p, _ := NewPool(urls, do)
	go p.Run()
	for i := 0; i < 7; i++ {
		chanIds <- strconv.Itoa(i)
	}
	time.Sleep(7 * time.Second)
}

func do(url string) error {
	for {
		select {
		case id := <-chanIds:
			switch url {
			case A:
				time.Sleep(2 * time.Second)
			case B:
				time.Sleep(3 * time.Second)
			}
			fmt.Println(fmt.Sprintf("do url, time:%s, url:%s, id:%s", time.Now().Format("2006-01-02 15:04:05"), url, id))
			return nil
		default:
			time.Sleep(time.Second)
		}
	}
}
