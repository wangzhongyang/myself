package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const fileFormName = "/Users/wangzhongyang/Documents/SpeakIn/voice/five_minute/3593.wav"

//const fileFormName = "/Users/wangzhongyang/Documents/SpeakIn/文档/语音质量检测/国家库-speakin.postman_collection.json"
const Url = "http://192.168.0.232:6004/check"

func main() {
	//f, err := ioutil.ReadFile(fileFormName)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(string(f))
	file, err := os.Open(fileFormName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	fmt.Println(len(b), "-----", b[100])
	req, _ := http.NewRequest("POST", Url, bytes.NewBuffer(b))

	req.Header.Add("Accept", "*/*")
	req.Header.Add("Accept-Encoding", "gzip, deflate")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Content-Type", "application/octet-stream")
	req.Header.Add("cache-control", "no-cache,no-cache")
	//req.Header.Add("Content-Length", strconv.Itoa(len(b)))

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res.StatusCode)
	fmt.Println(string(body))

}
