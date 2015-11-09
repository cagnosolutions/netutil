package netutil

import "net"

type NetHandler interface {
	HandleConn(c *net.Conn)
	//TODO: finish this package
}
