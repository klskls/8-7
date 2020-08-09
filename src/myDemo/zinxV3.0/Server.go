package main

import (
	"8-7/src/zinx/ziface"
	"8-7/src/zinx/znet"
	"fmt"
)

type MyRouter struct {
	znet.BaseRouter
}

func (b *MyRouter) PerHandle(request ziface.IRequest) {
	fmt.Println("per")
	request.GetConnection().GetTCPConnection().Write([]byte("pre"))
}

func (b *MyRouter) Handle(request ziface.IRequest) {
	fmt.Println("handler")
	request.GetConnection().GetTCPConnection().Write([]byte("handler"))
}

func (b *MyRouter) PostHandle(request ziface.IRequest) {
	fmt.Println("post")
	request.GetConnection().GetTCPConnection().Write([]byte("post"))
}

//GOARCH=amd64;GOOS=linux
func main() {

	s := znet.NewServer("zinx")
	s.AddRouter(&MyRouter{})
	s.Serve()
}
