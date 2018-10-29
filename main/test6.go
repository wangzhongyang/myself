package main

import (
	"fmt"
	"time"
)

func main() {
	//arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 9, 8, 7, 6, 5, 4, 3, 2, 10}
	//fmt.Println(RemoveRepByLoop(arr))
	var a int
	a = 1000000000000000000
	fmt.Println(a)
	arr := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(DelByValue(arr, 3))

	fmt.Println(time.Now().Format("2006-01-02T15:04:05Z07:00"))

}
func RemoveRepByLoop(slc []int) []int {
	result := make([]int, 0) // 存放结果
	for i := range slc {
		flag := true
		for j := range result {
			if slc[i] == result[j] {
				flag = false // 存在重复元素，标识为false
				break
			}
		}
		if flag { // 标识为false，不添加进结果
			result = append(result, slc[i])
		}
	}
	return result
}

func DelByValue(arr []int, val int) []int {
	for k, v := range arr {
		if v == val {
			arr = append(arr[:k], arr[k+1:]...)
			return arr
		}
	}
	return arr
}
