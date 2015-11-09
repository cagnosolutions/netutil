package netutil

import "net"

type NetServer struct {
	Conns    []*net.Conn
	Handlers []NetHandler
}

func NewNetServer() *NetServer {
	return &NetServer{
		Conns:    make([]*net.Conn, 0),
		Handlers: make([]NetHandler, 0),
	}
}

func (s *NetServer) Add(h NetHandler) {
	s.Handlers = append(s.Handlers, h)
}

func (s *NetServer) HandleConn(c *net.Conn) {
	return
}
