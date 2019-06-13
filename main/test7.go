package main

import (
	"fmt"
	"net/url"
)

type ProductGraph struct {
	ID        int  `json:"id"`
	ProductID int  `json:"product_id"`
	Default   bool `json:"default"`
}

func main() {
	ps := []ProductGraph{
		ProductGraph{ID: 1, ProductID: 1010513, Default: false},
		ProductGraph{ID: 2, ProductID: 1010513, Default: false},
		ProductGraph{ID: 3, ProductID: 1010513, Default: false},
		ProductGraph{ID: 4, ProductID: 1010513, Default: true},
		ProductGraph{ID: 5, ProductID: 1010513, Default: true},
	}
	a := A(ps)
	fmt.Println(a, a[1010513].ProductID, a[1010513].ID)

	encodeurl := "file:///storage/emulated/0/Android/data/com.bindo.stocktakeapp.beta/files/%E4%BD%A0%E5%A5%BD%20%20%E6%97%A9%E4%B8%8A%E5%A5%BD_181122105424.csv"
	decodeurl, err := url.QueryUnescape(encodeurl)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(decodeurl)
}

func A(ps []ProductGraph) map[int]ProductGraph {
	a := make(map[int]ProductGraph)
	for _, v := range ps {
		temp, ok := a[v.ProductID]
		if ok {
			vDefault := BoolToInt(v.Default)
			tempDefault := BoolToInt(temp.Default)
			if vDefault > tempDefault {
				a[v.ProductID] = v
			}
			if vDefault == tempDefault && v.ID < temp.ID {
				a[v.ProductID] = v
			}
		} else {
			a[v.ProductID] = v
		}
	}
	return a
}

func BoolToInt(b bool) int {
	if b {
		return 1
	} else {
		return 0
	}
}
