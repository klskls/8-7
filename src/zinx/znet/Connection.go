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

	Router ziface.IRouter

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
		re := Request{
			conn: c,
			data: buf[:cnt],
		}
		/*c.handleFunc(c.Conn, buf, cnt)*/
		go func(request *Request) {
			c.Router.PerHandle(request)
			c.Router.Handle(request)
			c.Router.PostHandle(request)
		}(&re)
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

func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

func NewConnection(conn *net.TCPConn, connID uint32, router ziface.IRouter) *Connection {
	c := &Connection{
		Conn:     conn,
		ConnID:   connID,
		Router:   router,
		IsClosed: false,
		ExitChan: make(chan bool, 1),
	}
	return c
}
