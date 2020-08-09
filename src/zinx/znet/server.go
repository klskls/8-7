package znet

import (
	"8-7/src/zinx/ziface"
	"fmt"
	"net"
)

type Server struct {
	Name       string
	IPVsersion string
	IP         string
	Port       int
}

func (s *Server) Start() {
	go func() {
		fmt.Println("[start] server")
		addr, err := net.ResolveTCPAddr(s.IPVsersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("error")
			return
		}
		lister, _ := net.ListenTCP(s.IPVsersion, addr)

		for {
			conn, _ := lister.AcceptTCP()
			go func() {
				for {
					buf := make([]byte, 512)
					cnt, _ := conn.Read(buf)
					conn.Write(buf[:cnt])
				}
			}()
		}
	}()

}

func (s *Server) Stop() {

}

func (s *Server) Serve() {
	s.Start()
	select {}
}

func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:       name,
		IPVsersion: "tcp4",
		IP:         "0.0.0.0",
		Port:       8999,
	}
	return s
}
