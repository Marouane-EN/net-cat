// Optional client mode
package client

import (
	"fmt"
	"net"
)

type Client struct {
	Name string
	Conn net.Conn
}

func GetClientName(conn net.Conn) string {
	var name string
	fmt.Fscanln(conn, &name)
	return name
}
