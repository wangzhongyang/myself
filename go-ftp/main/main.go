package main

import (
	"fmt"
	"log"
	"time"

	"github.com/jlaffaye/ftp"
)

func main() {
	c, err := ftp.Dial("127.0.0.1:30060", ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		log.Fatal(err)
	}

	err = c.Login("young", "12345")
	if err != nil {
		log.Fatal(err)
	}

	// Do something with the FTP conn
	list, err := c.List("sql")
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range list {
		fmt.Println(v.Name, v.Size)
	}
}
