package main

import (
	"fmt"
)

//实现f(n)函数，f(n）返回1-n之间整数不带7数字的个数；如17就是一个带7的数字，18是不带7的数字；
func main() {
	fmt.Println(f(4, 0, 1))
}

func f(n, n1, n2 int) int {
	if n == 0 {
		return n1
	}
	return f(n-1, n2, n1+n2)
}
