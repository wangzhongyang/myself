package main

import (
	"encoding/hex"
	"math/rand"

	"fmt"
	"strconv"
)

type ComplexSearchResult struct {
	ID      int    `json:"id"`
	StoreID int    `json:"store_id"`
	Name    string `json:"name"`
}

func main() {
	//	str := `
	//[{"id":631388,"store_id":176,"name":"Natural Balance Dry Cat Food Indoor Ultra Premium Formula 6Lb"},{"id":632556,"store_id":176,"name":"Untitled - 816587010772"},{"id":633497,"store_id":176,"name":"Honest Kitchen KEEN 2lb"},{"id":633504,"store_id":176,"name":"ACANA LI Lamb \u0026 Apple 4.4#"},{"id":634150,"store_id":176,"name":"Untitled - 667902164188"},{"id":634152,"store_id":176,"name":"Untitled - 667902172442"},{"id":598258,"store_id":5417,"name":"Blue Buffalo - Blue™ Wilderness™ Wild Bites™ Dog Trail Treat™ Salmon 4oz"},{"id":600337,"store_id":176,"name":"Four Paws® Wee Wee Pads® Puppy Housebreaking Pads 22\" X 23\" 10pads"},{"id":602045,"store_id":5417,"name":"Green Travel Cup - Small"},{"id":602863,"store_id":176,"name":"Old Mother Hubbard Classic Crunchy P-Nuttier Oven-Baked Dog Biscuits Mini 5Oz"}]
	//`
	//	var arr []ComplexSearchResult
	//	json.Unmarshal([]byte(str), &arr)
	//	fmt.Println(arr)
	//	m := make(map[int][]int)
	//	for _, v := range arr {
	//		m[v.StoreID] = append(m[v.StoreID], v.ID)
	//
	//	}
	//	fmt.Println(m)
	fmt.Println(strconv.Atoi("123455A"))
}

func randomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
