package znet

import (
	"8-7/src/zinx/ziface"
	"errors"
	"fmt"
	"net"
)

type Server struct {
	Name       string
	IPVsersion string
	IP         string
	Port       int
	Router     ziface.IRouter
}

func CallBackClient(conn *net.TCPConn, buf []byte, cnt int) error {
	fmt.Println("Conn Handle")
	if _, err := conn.Write(buf[:cnt]); nil != err {
		fmt.Println("write back buf err", err)
		return errors.New("CallBackToClient error")
	}
	return nil
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
		var cid uint32
		for {
			conn, _ := lister.AcceptTCP()
			dealConn := NewConnection(conn, cid, s.Router)
			cid++
			go dealConn.Start()
			/*go func() {
				for {
					buf := make([]byte, 512)
					cnt, _ := conn.Read(buf)
					conn.Write(buf[:cnt])
				}
			}()*/
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
		Router:     nil,
	}
	return s
}

func (s *Server) AddRouter(router ziface.IRouter) {
	s.Router = router
}
