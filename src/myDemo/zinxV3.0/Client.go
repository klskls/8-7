package main

import (
	"fmt"
	"net"
)

func main() {
	conn, _ := net.Dial("tcp", "127.0.0.1:8999")

	for {
		fmt.Println("write:asdadada")
		conn.Write([]byte("asdadada"))
		buf := make([]byte, 512)
		cnt, _ := conn.Read(buf)
		fmt.Println("reda:", buf[:cnt])

	}
}
