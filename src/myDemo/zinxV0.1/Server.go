package main

import "8-7/src/zinx/znet"

//GOARCH=amd64;GOOS=linux
func main() {
	s := znet.NewServer("zinx")
	s.Serve()
}
