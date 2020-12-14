package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"time"
)

// 省份校验
var vcity = []int{
	11,
	12,
	13,
	14,
	15,
	21,
	22,
	23,
	31,
	32,
	33,
	34,
	35,
	36,
	37,
	41,
	42,
	43,
	44,
	45,
	46,
	50,
	51,
	52,
	53,
	54,
	61,
	62,
	63,
	64,
	65,
	71,
	81,
	82,
	91,
}

// isCardNo 检查号码是否符合规范，包括长度，类型
func isCardNo(card string) bool {
	//身份证号码为15位或者18位，15位时全为数字，18位前17位为数字，最后一位是校验位，可能为数字或字符X
	matched, err := regexp.MatchString(`(^\d{15}$)|(^\d{17}(\d|X)$)`, card)
	if matched == false || err != nil {
		return false
	}
	return true
}

// checkProvince 取身份证前两位,校验省份
func checkProvince(card string) bool {
	province, err := strconv.Atoi(card[:2])
	if err != nil {
		return false
	}
	// 省份校验
	n := sort.SearchInts(vcity, province)
	if n >= 0 && n < len(vcity) {
		return true
	}
	return false
}

// verifyBirthday 校验日期
func verifyBirthday(year, month, day int, birthday time.Time) bool {
	now := time.Now()
	if birthday.Year() == year && birthday.Month() == time.Month(month) && birthday.Day() == day {
		//判断年份的范围（3岁到100岁之间)
		age := now.Year() - birthday.Year()
		if age >= 3 && age <= 100 {
			return true
		}
		return false
	}
	return false
}

// checkBirthday 检查生日是否正确
func checkBirthday(card string) bool {
	l := len(card)
	//身份证15位时，次序为省（3位）市（3位）年（2位）月（2位）日（2位）校验位（3位），皆为数字
	if l == 15 {
		year, _ := strconv.Atoi(card[6:8])
		year = 1900 + year
		month, _ := strconv.Atoi(card[8:10])
		day, _ := strconv.Atoi(card[10:12])
		birthday := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
		return verifyBirthday(year, month, day, birthday)
	}
	//身份证18位时，次序为省（3位）市（3位）年（4位）月（2位）日（2位）校验位（4位），校验位末尾可能为X
	if l == 18 {
		year, _ := strconv.Atoi(card[6:10])
		month, _ := strconv.Atoi(card[10:12])
		day, _ := strconv.Atoi(card[12:14])
		birthday := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
		return verifyBirthday(year, month, day, birthday)
	}
	return false
}

// changeFifteenToEighteen 15位转18位身份证号
func changeFifteenToEighteen(card string) string {
	if len(card) == 15 {
		arrInt := []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
		arrCh := []string{"1", "0", "X", "9", "8", "7", "6", "5", "4", "3", "2"}
		cardTemp := 0
		card = fmt.Sprintf("%s%s%s", card[0:6], "19", card[6:])
		for i := 0; i < 17; i++ {
			temp, _ := strconv.Atoi(card[i : i+1])
			cardTemp += temp * arrInt[i]
		}
		card += arrCh[cardTemp%11]
		return card
	}
	return card
}

//校验位的检测
func checkParity(card string) bool {
	//15位转18位
	card = changeFifteenToEighteen(card)
	l := len(card)
	if l == 18 {
		arrInt := []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
		arrCh := []string{"1", "0", "X", "9", "8", "7", "6", "5", "4", "3", "2"}
		cardTemp := 0
		for i := 0; i < 17; i++ {
			temp, _ := strconv.Atoi(card[i : i+1])
			cardTemp += temp * arrInt[i]
		}
		if arrCh[cardTemp%11] == card[17:18] {
			return true
		}
		return false
	}
	return false
}

func Verification(card string) bool {
	//校验长度，类型
	if !isCardNo(card) {
		return false
	}
	//检查省份
	if !checkProvince(card) {
		return false
	}
	//校验生日
	if !checkBirthday(card) {
		return false
	}
	//检验位的检测
	if !checkParity(card) {
		return false
	}
	return true
}
