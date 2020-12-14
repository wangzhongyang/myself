package main

import (
	"encoding/json"
	"fmt"
)

type NodeData struct {
	DeviceNum int64 `json:"device_num"`
	VoiceNum  int64 `json:"voice_num"`
}

func main() {
	m := make(map[int]NodeData)
	temp := m[9]
	temp.VoiceNum += 10
	temp.DeviceNum += 20
	m[9] = temp
	s, _ := json.Marshal(m)
	fmt.Println(string(s))
	var inspection ContentForSingleInspection
	fmt.Println(json.Unmarshal([]byte(inspectionStr), &inspection))
	fmt.Println(inspection)
	fmt.Println(json.Valid([]byte(inspectionStr)), json.Valid([]byte("{nihao a}")))
}

type ContentForSingleInspection struct {
	ChatList   []ContentForSingleInspectionChatItem `json:"chat_list"`
	Supplement string                               `json:"supplement"`
	Images     []string                             `json:"images"`
}

type ContentForSingleInspectionChatItem struct {
	MsgID           string `json:"msg_id"`
	UserAssistantID int    `json:"user_assistant_id"`
	UserName        string `json:"user_name"`
	IsCustomer      bool   `json:"is_customer"`
	MsgTime         int64  `json:"msg_time"` // 10位时间戳
	MsgType         int    `json:"msg_type"` // 1:text 2:link 3:image 4:emoji 5:unknown
	MsgContent      string `json:"msg_content"`
}

func (c *ContentForSingleInspection) UnmarshalJSON(b []byte) error {
	c1 := struct {
		ChatList   []ContentForSingleInspectionChatItem `json:"chat_list"`
		Supplement string                               `json:"supplement"`
		Images     []string                             `json:"images"`
	}{}
	if err := json.Unmarshal(b, &c1); err != nil {
		return err
	}
	c.Images = c1.Images
	c.ChatList = c1.ChatList
	c.Supplement = c1.Supplement
	return nil
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
