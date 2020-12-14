package main

import (
	"encoding/base64"
	"fmt"
	"git.forchange.cn/framework/utils"
	"reflect"
)

func main() {
	fmt.Println(reflect.TypeOf(float32(1)/float32(3)).Kind().String(), float32(1)/float32(3))
}

func decodeMain() {
	decode("OQXcneG-zMoregnz2iJgKSjZrEVXoP7Lgn1cIXJhPRXyFFK2yswWY2j2vfJv3Cm4Y8GvntW7F7PmELiFf18x1jekTBvBHa04vePwCaLX2u8=")
	decode("OQXcneG-zMoregnz2iJgKdiqw1Qysz2UbEq7um-71fQOjS6E4ITpl64xYGUn41FTQoZsZ5uEWk81oA_f4hKx0ey6pm3ylzJjLVDnnsNtUEpBDO0Maj6lAw==")
	decode("OQXcneG-zMoJPvjRSJzQo9mMk9Ce-0UcqtdhkNvtp1G0qBQEaCDFOHfwbfhtfc5MA33WsOeU7kV_yxrXtt2jSsgHIQxFoyWMvi_y2kq-QQA79U9iruzmkXUajti7xElJ1kTGhsm4HbM=")

}

const DesKey = "Fc@$#69^"

func decode(uaf string) {
	ds, _ := base64.URLEncoding.DecodeString(uaf)
	data, er := utils.Decrypt(ds, []byte(DesKey))
	if er != nil {
		panic("解密configuration.uaf出错")
	}
	fmt.Println("decode:", string(data))
}
