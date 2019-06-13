package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type TranslationInfo map[string]TranslationInfoItem
type TranslationInfoItem map[string]string

type WangMap map[string]string

type YangMap map[string]string

type Wang struct {
	Name   string  `json:"name"`
	Config WangMap `json:"config"`
}

type Yang struct {
	LastName string  `json:"last_name"`
	Config   YangMap `json:"config"`
}

func main() {
	yang := Yang{
		LastName: "wangzhongyang",
		Config: map[string]string{
			"1": "2",
			"3": "4",
		},
	}

	//field, _ := reflect.TypeOf(yang).FieldByName("LastName")
	s := reflect.ValueOf(&yang).Elem()
	a := map[string]string{
		"5": "6",
		"7": "8",
		"9": "10",
	}
	fmt.Printf("\nyang:        %+v\n", yang)
	//index := reflect.ValueOf(yang).FieldByName("LastName")
	s.FieldByName("Config").Set(reflect.ValueOf(a))
	//setConfig(yang, "LastName", "last------Name")
	fmt.Printf("\nyang:        %+v\n", yang)

	//translate_info: {
	//"field_name" :
	//
	//{ "zh_cn" : "xxx", "en" : "xxx" }
	//}
	translation := TranslationInfo{
		"field": TranslationInfoItem{
			"zh_cn": "xxx",
			"en":    "xxx",
		},
		"field2": TranslationInfoItem{
			"zh_cn": "xxx",
			"en":    "xxx",
		},
	}
	str, _ := json.Marshal(translation)
	fmt.Println(string(str))
}

func setConfig(data interface{}, fieldName string, fieldValue interface{}) error {
	//getType := reflect.TypeOf(data)
	fmt.Println("after:		", data)
	getValue := reflect.ValueOf(&data).Elem()
	getValue.FieldByName(fieldName).SetString(fieldValue.(string))

	return nil
}
