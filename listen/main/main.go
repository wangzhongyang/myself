package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

func handleConnection(conn net.Conn) {
	str := exec_shell(GetCurrentDirectory())
	setLog(str)
	conn.Close()
}

func main() {
	//ln, err := net.Listen("tcp", ":18080")
	//if err != nil {
	//	panic(err)
	//}
	//for {
	//	conn, err := ln.Accept()
	//	if err != nil {
	//		log.Fatal("get client connection error: ", err)
	//	}
	//	fmt.Println("------------------")
	//	handleConnection(conn)
	//}

	http.HandleFunc("/", foo)
	http.ListenAndServe(":18080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Server", "A Go Web Server")
	w.WriteHeader(http.StatusOK)
	str := `{
  "success":true
}
`
	w.Write([]byte(str))
	go build()
}

func build() {
	fmt.Println("-------build--------")
	res := exec_shell(GetCurrentDirectory())
	setLog(res)
}

func exec_shell(path string) string {
	cmd := exec.Command(path + "/" + "restart.sh") //初始化Cmd
	out, err := cmd.Output()                       //运行脚本
	if nil != err {
		fmt.Println("cmd output:		", err.Error())
	}
	err = cmd.Wait() //等待执行完成
	if nil != err {
		fmt.Println("cmd wait:		", err.Error())
	}
	return string(out)
}
func setLog(str string) {
	// 定义一个文件
	fileName := "restart-goVps.log"
	// 文件名，只写|追加|创建，权限
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("can't open file")
	}
	defer f.Close()
	str = fmt.Sprintf("-----------------%s--------------------\n%s", time.Now().Format("2006-01-02 15:04:05"), str)
	_, _ = f.WriteString(str)
}
func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0])) //返回绝对路径  filePath.Dir(os.Args[0])去除最后一个元素的路径
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1) //将\替换成/
}
