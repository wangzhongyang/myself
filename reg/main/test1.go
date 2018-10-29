package main

import (
	"fmt"
	"regexp"
)

func main() {
	// 匹配小数
	str := "<![CDATA[10.99%]]>"
	regStr := `[0-9]+(\.[0-9]{1,3})?`
	reg := regexp.MustCompile(regStr)
	fmt.Println(reg.FindString(str))
}
