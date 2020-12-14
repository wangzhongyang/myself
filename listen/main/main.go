package main

import (
	"fmt"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic("listen tcp failed," + err.Error())
	}
	go func() {
		for {
			conn, err := ln.Accept()
			if err != nil {
				panic("conn failed TCP," + err.Error())
			}
			_, _ = conn.Write([]byte("this is TCP"))
			_ = conn.Close()
		}
	}()

	ln2, err := net.ListenPacket("udp", ":8080")
	if err != nil {
		panic("listen UDP failed," + err.Error())
	}
	defer ln2.Close()
	for {
		buf := make([]byte, 1024)
		_, addr, err := ln2.ReadFrom(buf)
		if err != nil {
			continue
		}
		fmt.Println(addr.String())
		if _, err := ln2.WriteTo([]byte("this is UDP\n"), addr); err != nil {
			fmt.Println("udp write failed," + err.Error())
		}
	}
}

func handleConnectionTcp(conn net.Conn) {
	_, _ = conn.Write([]byte("this is TCP"))
	_ = conn.Close()
}
