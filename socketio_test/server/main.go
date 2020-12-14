package main

import (
	"fmt"
)

func main() {
	fmt.Println(returnValues())
}

func returnValues() int {
	var result int
	defer func() {
		result++
		fmt.Println("defer")
	}()
	return result
}
