package znet

import (
	"8-7/src/zinx/ziface"
	"fmt"
	"net"
)

type Connection struct {
	Conn *net.TCPConn

	ConnID uint32

	IsClosed bool

	handleFunc ziface.HandleFunc

	ExitChan chan bool
}

func (c *Connection) Start() {
	fmt.Println("Conn start()...")
	go c.StartReader()
}

func (c *Connection) StartReader() {
	fmt.Println("Reader is running")
	defer fmt.Println("Reader is Exit")
	defer c.Stop()

	for {
		buf := make([]byte, 512)
		cnt, _ := c.Conn.Read(buf)
		fmt.Println(buf[:cnt])
		c.handleFunc(c.Conn, buf, cnt)
	}
}

func (c *Connection) Stop() {
	fmt.Println("Conn stop()。。。。")
	if c.IsClosed == true {
		return
	}
	c.IsClosed = true

	c.Conn.Close()

	close(c.ExitChan)
}

func NewConnection(conn *net.TCPConn, connID uint32, callBack_api ziface.HandleFunc) *Connection {
	c := &Connection{
		Conn:       conn,
		ConnID:     connID,
		handleFunc: callBack_api,
		IsClosed:   false,
		ExitChan:   make(chan bool, 1),
	}
	return c
}
