package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	a, b := []int{1, 2, 3, 4}, []int{3, 4, 5, 6, 7}
	fmt.Println(Different(a, b))
	c := make(map[string]int)
	c["1"] = 1
	c["2"] = 2
	fmt.Println(len(c))
	key := "4036263823jVSfRiRQeDUVTyGdCnbcBA=="
	fmt.Println(key[10:])

	str := `{"sid":10001,"oid":10002,"n":"20180010000001" ,"pa":"10.02","et":1380000000, "of": 1}`
	var mq MqMessage
	fmt.Println(json.Unmarshal([]byte(str), &mq))
	fmt.Printf("mq:		\n%+v\n", mq)
}

type MqMessage struct {
	Sid        int    `json:"sid"`
	Oid        int    `json:"oid"`
	N          string `json:"n"`
	Et         int64  `json:"et"`
	PaidAmount string `json:"pa"`
	PushMethod int    `json:"push_method"`
}

func Different(a, b []int) ([]int, []int) {
	res1, res2 := make([]int, 0), make([]int, 0)
	for _, v := range a {
		key := Search(b, v)
		if len(b) == key { // 不能查找
			res1 = append(res1, v)
		} else {
			b = append(b[:key], b[key+1:]...)
		}
	}
	res2 = b
	return res1, res2
}

func Search(a []int, b int) int {
	for k, v := range a {
		if v == b {
			return k
		}
	}
	return len(a)
}
