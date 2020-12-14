package main

import (
	"fmt"
	"net/http"
	"regexp"
)

func main() {
	http.HandleFunc("/", route)
	http.ListenAndServe(":8080", nil)
}

var num = regexp.MustCompile(`\d`)
var str1 = regexp.MustCompile(`\w`)

func route(w http.ResponseWriter, r *http.Request) {
	switch {
	case num.MatchString(r.URL.Path):
		fmt.Println("\\d....")
		digits(w, r)
	case str1.MatchString(r.URL.Path):
		fmt.Println("\\w....")
		str(w, r)
	default:
		w.Write([]byte("位置匹配项"))
	}
}

func digits(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("receive digits"))
}

func str(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("receive string"))
}
