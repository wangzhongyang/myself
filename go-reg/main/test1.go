package main

import (
	"encoding/json"
	"fmt"
)

type OrderCreateMqMessage struct {
	StoreID     int    `json:"store_id"`
	StoreSlug   string `json:"store_slug"`
	OrderID     int    `json:"order_id"`
	OrderNumber string `json:"order_number"`
	OrderFrom   int    `json:"order_from"`
	I18nLang    string `json:"i18n_lang"`
	Event       int    `json:"event"`
}

func main() {
	// 长度相同用后面的

	str := "abcdasdfghjkldcba"
	fmt.Println(longestPalindrome(str))

	var temp OrderCreateMqMessage
	str1 := `{"store_id":6994, "store_slug":"2b6", "order_id":3568024, "order_number":"201812201431175848084878", "order_from":4, "i18n_lang":"zh_hans_cn", "event":0}`
	if err := json.Unmarshal([]byte(str1), &temp); err != nil {
		fmt.Println(err.Error())
	} else {

		fmt.Printf("\ntemp:        %+v\n", temp)
	}
}

func longestPalindrome(s string) string {
	return ""
}
