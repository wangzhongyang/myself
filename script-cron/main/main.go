package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type ScriptCronTimeType string

const (
	ScriptCronTimeTypeDay    = "day"
	ScriptCronTimeTypeHour   = "hour"
	ScriptCronTimeTypeMinute = "minute"
	ScriptCronTimeTypeSecond = "second"
)

type ScriptCron struct {
	NowTime   time.Time `json:"now_time"`
	NowDay    int       `json:"now_day"`
	NowHour   int       `json:"now_hour"`
	NowMinute int       `json:"now_minute"`
	NowSecond int       `json:"now_second"`
}

func main() {
	// 分 时 天
	str := "10-20 * * *"
	fmt.Println("begin:		", time.Now().Format("15:04:05"))
	for {
		s := Init()
		if temp, _ := s.IsRunningTime(str); temp {
			fmt.Println("IsRunningTime:		", time.Now().Format("15:04:05"))
		}
		fmt.Println("for:		", s.NowTime.Format("15:04:05"))
		time.Sleep(time.Second)
	}
}

func Init() *ScriptCron {
	s := new(ScriptCron)
	s.NowTime = time.Now()
	s.NowDay, s.NowHour, s.NowMinute, s.NowSecond = s.NowTime.Day(), s.NowTime.Hour(), s.NowTime.Minute(), s.NowTime.Second()
	return s
}

func (s *ScriptCron) IsRunningTime(cronStr string) (bool, error) {
	timeTypeMap := map[int]ScriptCronTimeType{
		0: ScriptCronTimeTypeSecond,
		1: ScriptCronTimeTypeMinute,
		2: ScriptCronTimeTypeHour,
		3: ScriptCronTimeTypeDay,
	}
	cronArray := strings.Split(cronStr, " ")
	cronArrayLength := len(cronArray)
	for timeTypeNumber, cronStrItem := range cronArray {
		if cronArrayLength == 3 {
			timeTypeNumber++
		}
		isMatch, err := s.IsMatch(cronStrItem, timeTypeMap[timeTypeNumber])
		if err != nil {
			return false, errors.New(fmt.Sprintf("range cronArray error,error:%v", err))
		}
		if !isMatch {
			return false, nil
		}
	}
	return true, nil
}

func (s *ScriptCron) IsMatch(cronStr string, timeType ScriptCronTimeType) (bool, error) {
	timeNumber := s.GetTimeNumber(timeType)
	if cronStr == "*" {
		return true, nil
	}

	if strings.Contains(cronStr, "/") { // */5 当前粒度与5取余是否为0 ，成立为true
		temp := strings.Split(cronStr, "/")
		if len(temp) == 2 {
			divisor, err := strconv.Atoi(temp[1])
			if err != nil {
				return false, errors.New(fmt.Sprintf("cronStr '%s' cant format to int,error:%v", cronStr, err))
			}
			fmt.Println(timeNumber, divisor, timeNumber%divisor)
			if timeNumber%divisor == 0 {
				return true, nil
			}
		}
	}

	// cronStr := "1,2,3,4,5,15"
	if tempArray := strings.Split(cronStr, ","); len(tempArray) > 1 { // 存在以 ‘,’ 分割的数，
		timeNumberStr := strconv.Itoa(timeNumber)
		for _, v := range tempArray {
			if v == timeNumberStr {
				return true, nil
			}
		}
	}

	// cronStr := "1-5"
	if tempArray := strings.Split(cronStr, "-"); len(tempArray) > 1 { // 存在以 ‘-’ 分割的数，表示范围
		begin, _ := strconv.Atoi(tempArray[0])
		end, _ := strconv.Atoi(tempArray[1])
		for begin <= end {
			if begin == timeNumber {
				return true, nil
			}
			begin++
		}
	}
	return false, nil
}

func (s *ScriptCron) GetTimeNumber(timeType ScriptCronTimeType) int {
	switch timeType {
	case ScriptCronTimeTypeSecond:
		return s.NowSecond
	case ScriptCronTimeTypeMinute:
		return s.NowMinute
	case ScriptCronTimeTypeHour:
		return s.NowHour
	case ScriptCronTimeTypeDay:
		return s.NowDay
	}
	return 0
}
