package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"sync"
)

//AI核心代码，估值一个亿
func OneHundredMillion() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		s1 := input.Text()
		s1 = strings.Replace(s1, "吗", "", -1)
		s1 = strings.Replace(s1, "?", "！", -1)
		s1 = strings.Replace(s1, "？", "！", -1)
		fmt.Println(s1)
	}
}

const inspectionStr = `{
    "chat_list": [
        {
            "msg_id": "odijgjjfndjdjfjghghg",
            "user_assistant_id": 213231,
            "user_name": "用户名字",
            "is_customer": true,
            "msg_time": 1234567890,
            "msg_type": 1,
            "msg_content": "消息内容1"
        },
        {
            "msg_id": "odijgjjfndjdjfjghghg",
            "user_assistant_id": 213231,
            "user_name": "用户名字",
            "is_customer": true,
            "msg_time": 1234567890,
            "msg_type": 1,
            "msg_content": "消息内容2"
        },
        {
            "msg_id": "odijgjjfndjdjfjghghg",
            "user_assistant_id": 213231,
            "user_name": "用户名字",
            "is_customer": true,
            "msg_time": 1234567890,
            "msg_type": 1,
            "msg_content": "消息内容3"
        },
        {
            "msg_id": "odijgjjfndjdjfjghghg",
            "user_assistant_id": 213231,
            "user_name": "用户名字",
            "is_customer": false,
            "msg_time": 1234567890,
            "msg_type": 1,
            "msg_content": "消息内容4"
        }
    ],
    "supplement": "补充说明",
    "images": [
        "https://thirdwx.qlogo.cn/mmopen/vi_32/V5rIrvQvpibIg6V3Ja3ELmLhce0icfWyhEF2pkSUhHOzqwtg5DoyqPHribwbvS4fPrEfIOCoGLvnNEZBcia1MChQUA/132",
        "https://thirdwx.qlogo.cn/mmopen/vi_32/V5rIrvQvpibIg6V3Ja3ELmLhce0icfWyhEF2pkSUhHOzqwtg5DoyqPHribwbvS4fPrEfIOCoGLvnNEZBcia1MChQUA/132"
    ]
}`

type String1 int

const (
	Name String1 = iota
	Age
)

func (s String1) String() string {
	switch s {
	case Name:
		return "name"
	case Age:
		return "age"
	}
	return fmt.Sprintf("%d", s)
}

func main() {
	s := `{"age":3,"name":"name2"}`
	var p People
	fmt.Println(json.Unmarshal([]byte(s), &p))

}

type People struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (p *People) String() string {
	p.Name = "name2"
	return fmt.Sprintf("%s%d", p.Name, p.Age)
}

func (p *People) UnmarshalJSON(b []byte) error {
	fmt.Println("this is p UnmarshalJSON")
	p1 := struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{}
	if err := json.Unmarshal(b, &p1); err != nil {
		return err
	}
	p.Age = p1.Age
	p.Name = p1.Name
	return nil
}

var i int

func ToString(b *[]byte) string {
	i += 1
	if len(*b) == 0 {
		*b = []byte(fmt.Sprintf("%s%d", "hahahahahahhahahahahhahahahahahhahahahaahhahah", i))
	} else {
		*b = append(*b, byte(i))
	}

	return string(*b)
}

var PeoplePool = sync.Pool{New: func() interface{} {
	return &People{
		Name: "name1",
		Age:  3,
	}
}}

//func init() {
//	PeoplePool = sync.Pool{New: func() interface{} {
//		return make([]byte, 0)
//	}}
//}
func NewByPool() *People {
	return PeoplePool.Get().(*People)
}

func NeyPeople() *People {
	return &People{
		Name: "name1",
		Age:  3,
	}
}
