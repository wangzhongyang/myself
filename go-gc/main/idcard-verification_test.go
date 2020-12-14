package main

import (
	"fmt"
	"testing"
)

func Test_Void(t *testing.T) {
	// 一下输出应全为true
	//card := "360721199906013354"
	card := "360721990601335"
	fmt.Println(isCardNo(card))
	fmt.Println(checkProvince(card))
	fmt.Println(checkBirthday(card))
	fmt.Println(changeFifteenToEighteen(card) == "360721199906013354")
	fmt.Println(Verification(card))
}
