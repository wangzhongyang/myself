package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"unicode/utf8"

	"unicode/utf16"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

func main() {
	src := "编码转换内容内容"
	reader := transform.NewReader(bytes.NewReader([]byte(src)), simplifiedchinese.GBK.NewEncoder())
	data, _ := ioutil.ReadAll(reader)
	fmt.Println("void:		", utf8.Valid(data))
	fmt.Println("void:		", string(data))

	fmt.Println("------------------------------")
	fmt.Printf("%t, ", utf16.IsSurrogate(0xD400)) // false
	fmt.Printf("%t, ", utf16.IsSurrogate(0xDC00)) // true
	fmt.Printf("%t\n", utf16.IsSurrogate(0xDFFF)) // true

	r1, r2 := utf16.EncodeRune('𠀾')
	fmt.Printf("%x, %x\n", r1, r2) // d840, dc3e

	r := utf16.DecodeRune(0xD840, 0xDC3E)
	fmt.Printf("%c\n", r) // d840, dc3e

	u := []uint16{'不', '会', 0xD840, 0xDC3E}
	s := utf16.Decode(u)
	fmt.Printf("%c", s) // [不 会 𠀾]
}

func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0])) //返回绝对路径  filePath.Dir(os.Args[0])去除最后一个元素的路径
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1) //将\替换成/
}
