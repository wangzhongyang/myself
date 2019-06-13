package main

import (
	"errors"
	"fmt"
	"reflect"
)

type People struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	people1 := People{ID: 1, Name: "name1", Age: 11}
	fmt.Println(GetFieldValue("Name", people1))
	SetFieldValue("Name", people1)
	fmt.Println(people1.Name)

	type t struct {
		N int
	}
	var n = t{42}
	fmt.Println(n.N)
	reflect.ValueOf(&n).Elem().FieldByName("N").SetInt(7)
	fmt.Println(n.N)
	//people2 := People{ID: 2, Name: "name2", Age: 12}
	//peoples := []People{people1, people2}
	//res, err := GetFieldMap("ID", &peoples)
	//fmt.Println(res, err)
	//for k, v := range res {
	//	item := v.(People)
	//	str, _ := json.Marshal(item)
	//	fmt.Println("key,   value:		", k, string(str))
	//}
	//resArray, err := GetFieldArray("Age", &peoples)
	//fmt.Println(resArray, err)
}

func GetFieldValue(fieldName string, data interface{}) (string, error) {
	return reflect.ValueOf(data).FieldByName(fieldName).String(), nil
}

func SetFieldValue(fieldName string, data People) {
	reflect.ValueOf(&data).Elem().FieldByName(fieldName).Set(reflect.ValueOf("99999"))
	fmt.Println("reflect.ValueOf(&data).Elem():		", reflect.ValueOf(&data).Elem().Kind().String())

	fmt.Println("at SetFieldValue:", data.Name)

}

func GetFieldMap(fieldName string, data interface{}) (map[int]interface{}, error) {
	res := make(map[int]interface{})
	rt := reflect.TypeOf(data)

	if rt.Kind() != reflect.Ptr {
		return res, errors.New("can't set non ptr data")
	}
	if rt.Elem().Kind() != reflect.Slice && rt.Elem().Kind() != reflect.Array {
		return res, errors.New("can't set non slice data")
	}
	rv := reflect.ValueOf(data).Elem()
	lens := rv.Len()
	for i := 0; i < lens; i++ {
		row := rv.Index(i)
		if row.Type().Kind() != reflect.Struct {
			return res, errors.New("can't set non struct data on list")
		}
		if !row.FieldByName(fieldName).IsValid() {
			return nil, errors.New("cant get field by name:" + fieldName)
		}
		fieldValue := int(row.FieldByName(fieldName).Int())
		res[fieldValue] = row.Interface()
	}
	return res, nil
}

func GetFieldArray(fieldName string, data interface{}) ([]int, error) {
	res := make([]int, 0)
	rt := reflect.TypeOf(data)

	if rt.Kind() != reflect.Ptr {
		return res, errors.New("can't set non ptr data")
	}
	if rt.Elem().Kind() != reflect.Slice && rt.Elem().Kind() != reflect.Array {
		return res, errors.New("can't set non slice data")
	}
	rv := reflect.ValueOf(data).Elem()
	lens := rv.Len()
	for i := 0; i < lens; i++ {
		row := rv.Index(i)
		if row.Type().Kind() != reflect.Struct {
			return res, errors.New("can't set non struct data on list")
		}
		if !row.FieldByName(fieldName).IsValid() {
			return nil, errors.New("cant get field by name:" + fieldName)
		}
		fieldValue := int(row.FieldByName(fieldName).Int())
		res = append(res, fieldValue)
	}
	return res, nil
}
