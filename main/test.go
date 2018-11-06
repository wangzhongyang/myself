package main

import (
	"encoding/hex"
	"fmt"
	"math/rand"
	"regexp"
	"sort"
	"strings"

	"github.com/Pallinder/go-randomdata"
)

type ComplexSearchResult struct {
	ID      int    `json:"id"`
	StoreID int    `json:"store_id"`
	Name    string `json:"name"`
}

type Wang []int

func (s Wang) Len() int           { return len(s) }
func (s Wang) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s Wang) Less(i, j int) bool { return s[i] < s[j] }

func main() {
	details(" a o nnn w")
	a := Wang{1, 2, 44, 5, 77, 55, 31}
	sort.Sort(a)
	fmt.Println(a)

	for i := 0; i < 20; i++ {

		str := randomdata.Letters(24)
		fmt.Println(str, "========]", len(str))
	}

	fmt.Println("\uFFFD")

	bytes := make([]byte, 2014)
	fmt.Println(len(bytes), cap(bytes), string(bytes[0]))

	str := "王忠洋aaa"
	fmt.Println("汉子长度：		", strings.Count(str, "")-1)
	/***
	   sort.Sort(a)
	   notification.SendToDevice(a)
	   NewPlatform() Pltform {
	}
	*/

	SayHello("this is message")
}

func randomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func details(str string) string {
	str = regexp.MustCompile(" ").ReplaceAllString(str, "+")
	fmt.Println(str)
	return str
}

func SayHello(msg string) error {
	fmt.Println("send message is :", msg)
	return nil
}
