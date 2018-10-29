package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	layout := "2006-01-02 15:04:05"
	// Add 时间相加
	now := time.Now()
	// ParseDuration parses a duration string.
	// A duration string is a possibly signed sequence of decimal numbers,
	// each with optional fraction and a unit suffix,
	// such as "300ms", "-1.5h" or "2h45m".
	//  Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
	// 10分钟前
	m, _ := time.ParseDuration("-1m")
	m1 := now.Add(m)
	fmt.Println(m1)

	// 8个小时前
	h, _ := time.ParseDuration("-1h")
	h1 := now.Add(8 * h)
	fmt.Println(h1)

	// 一天前
	d, _ := time.ParseDuration("-24h")
	d1 := now.Add(d)
	fmt.Println(d1)

	printSplit(50)

	// 10分钟后
	mm, _ := time.ParseDuration("1m")
	mm1 := now.Add(mm)
	fmt.Println(mm1)

	// 8小时后
	hh, _ := time.ParseDuration("1h")
	hh1 := now.Add(hh)
	fmt.Println(hh1)

	// 一天后
	dd, _ := time.ParseDuration("24h")
	dd1 := now.Add(dd)
	fmt.Println(dd1)

	printSplit(50)

	// Sub 计算两个时间差
	subM := now.Sub(m1)
	fmt.Println(subM.Minutes(), "分钟")

	sumH := now.Sub(h1)
	fmt.Println(sumH.Hours(), "小时")

	sumD := now.Sub(d1)
	fmt.Printf("%v 天\n", sumD.Hours()/24)

	timeStr1 := "2018-01-01 15:04:05"
	timeStr2 := "2018-01-01 15:04:06"
	time1, _ := time.Parse(layout, timeStr1)
	time2, _ := time.Parse(layout, timeStr2)
	fmt.Println(timeStr1, timeStr2, "分钟差距：		", time1.Sub(time2).Minutes() < 0)

	// 时间戳转time
	nowUnix := time.Now().Unix()
	fmt.Println("时间戳转time：		", nowUnix, time.Unix(nowUnix, 0).Format(layout))
	time.Sleep(1 * time.Second)
	var nowNIl *time.Time
	nowNIl = nil
	fmt.Println(getTimeDifference(*nowNIl, -10))
	fmt.Println(voidTime(time.Now(), -10, 3))
}

func printSplit(count int) {
	fmt.Println(strings.Repeat("#", count))
}

// 计算给定时间的 前/后 分钟的距现在的毫秒数
func getTimeDifference(appointedTime time.Time, difference int) (int64, error) {
	//if appointedTime {
	//	return 99999999, nil
	//}
	if &appointedTime == nil {
		return 999999, nil
	}
	layout := "2006-01-02 15:04:05"
	timeStr1 := "2018-01-01 15:04:05" // 当前时间
	timeStr2 := "2018-01-01 15:34:06" // 指定时间
	time1, _ := time.Parse(layout, timeStr1)
	time2, _ := time.Parse(layout, timeStr2)

	appointedTime2 := time2.Add(time.Duration(difference) * time.Minute)
	sub := appointedTime2.Sub(time1).Nanoseconds() / 1000000
	return sub, nil
}

// 验证时间是否在 指定时间 的
func voidTime(appointedTime time.Time, difference, sub int) bool {
	layout := "2006-01-02 15:04:05"
	timeStr1 := "2018-01-01 15:04:05" // 当前时间
	timeStr2 := "2018-01-01 15:14:05" // 指定时间
	time1, _ := time.Parse(layout, timeStr1)
	time2, _ := time.Parse(layout, timeStr2)
	appointedTime2 := time2.Add(time.Duration(difference) * time.Minute)
	fmt.Println("voidTime appointedTime2:		", appointedTime2.Format(layout))
	fmt.Println("voidTime appointedTime2 sub:		", appointedTime2.Sub(time1).Minutes())
	subMinutes := appointedTime2.Sub(time1).Minutes()
	return subMinutes < float64(sub) && subMinutes > float64(-1*sub)
}
