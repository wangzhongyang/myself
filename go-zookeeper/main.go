package main

import (
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"log"
	"sync"
	"time"
)

func main() {
	//serviceNamingMain()
	thisTime, _ := time.ParseInLocation("2006-01-02 15:04:05", "2019-06-16 11:18:22", time.Local)
	fmt.Println(thisTime.Format("2006-01-02 15:04:05"))
}

// serviceNamingMain 服务命名测试
func serviceNamingMain() {
	conn, _, err := zk.Connect([]string{"49.235.114.97:2181"}, time.Second*30)
	if err != nil {
		fmt.Println("conn error:", err.Error())
		return
	}
	defer conn.Close()
	path := "/zk-book/server-naming-"
	name := serviceNamingGenerate(conn, path)
	fmt.Println("name:	", name)
}

func serviceNaming(conn *zk.Conn, path string) {

}

// serviceNamingGenerate 服务命名生成
func serviceNamingGenerate(conn *zk.Conn, path string) string {
	name, err := conn.CreateProtectedEphemeralSequential(path, []byte("1"), zk.WorldACL(0x1f))
	if err != nil {
		log.Fatalln("serviceNamingGenerate:", err.Error())
	}
	return name
}

// dataSource 配置中心示例
func dataSource() {
	conn, _, err := zk.Connect([]string{"49.235.114.97:2181"}, time.Second*30)
	if err != nil {
		fmt.Println("conn error:", err.Error())
		return
	}
	defer conn.Close()
	path := "/zk-book/data-source"
	for {
		source, stat, event, err := conn.GetW(path)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("source:", string(source), "stat:", stat.Version)
		select {
		case m := <-event:
			fmt.Println(fmt.Sprintf("event mgs:%+v", m))
			continue
		}
	}
}

// lock 分布式锁示例
func lock() {
	conn, _, err := zk.Connect([]string{"49.235.114.97:2181"}, time.Second*30)
	if err != nil {
		fmt.Println("conn error:", err.Error())
		return
	}
	defer conn.Close()
	path, _ := "/zk-book/lock-", new(sync.Once)
	lock := zk.NewLock(conn, path, zk.WorldACL(0x1f))
	num := 0
	for i := 0; i < 10; i++ {
		go func(n int) {
			if err := lock.Lock(); err != nil {
				fmt.Println("lock error:", n, err.Error())
				return
			}
			fmt.Println("number: ", num)
			num++
			_ = lock.Unlock()
		}(i)
	}
	time.Sleep(2 * time.Second)
}
