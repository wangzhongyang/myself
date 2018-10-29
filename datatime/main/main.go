package main

import (
	"fmt"
	"time"
)

func main() {
	// 格式化输出
	timeNowString := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println("timeNowString:	", timeNowString)
	// string to time
	stringToTime, _ := time.ParseInLocation("2006-01-02 15:04:05", "2018-09-09 15:20:37", time.Local)
	fmt.Println("string to time:		", stringToTime.Format("2006-01-02 15:04:05"))

	// 时间差
	sumH := time.Now().Sub(stringToTime)
	var a int
	a = 48
	fmt.Println("时间小时差：	", sumH.Hours() > float64(a))

	now := time.Now()
	stringToTime, _ = time.ParseInLocation("2006-01-02 15:04:05", now.Format("2006-01-02 15:04:05"), time.Local)
	fmt.Println(now.Sub(stringToTime).Hours() < 48.0)

	// 2小时前
	h, _ := time.ParseDuration("-2h")
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), time.Now().Add(h).Format("2006-01-02 15:04:05"))

	//
}
