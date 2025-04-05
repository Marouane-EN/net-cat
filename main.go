package main

import (
	"fmt"
	"log"
	"net"
)

type Server struct {
	listenAddr string
	ln         net.Listener
	quitch     chan struct{}
	msgch      chan []byte
}

func NewServer(addr string) *Server {
	fmt.Println("1")
	return &Server{
		listenAddr: addr,
		quitch:     make(chan struct{}),
		msgch:      make(chan []byte, 10),
	}
}

func (s *Server) Start() error {
	fmt.Println("2")
	ln, err := net.Listen("tcp", s.listenAddr)
	if err != nil {
		return err
	}
	defer ln.Close()
	s.ln = ln

	go s.acceptLoop()

	<-s.quitch
	close(s.msgch)
	return nil
}

func (s *Server) acceptLoop() {
	fmt.Println("3")
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			continue
		}
		conn.Write([]byte("[ENTER YOUR NAME]: "))
		
		go s.handleConn(conn)
	}
}

func (s *Server) handleConn(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 2048)
	validName := true
	var name string
	fmt.Println("4")
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read error:", err)
			return
		}
		s.msgch <- buf[:n]
		if validName {
			name = "[" + string(buf[:n-1]) + "]: "
			validName = false
		}
		
		conn.Write([]byte(name))

		// go s.writeLoop(conn)
	}
}

// func (s *Server) writeLoop(conn net.Conn) {
// 	for {
// 		select {
// 		case msg, ok := <-s.msgch:
// 			if !ok {
// 				return
// 			}
// 			_, err := conn.Write(msg)
// 			if err != nil {
// 				fmt.Println("write error:", err)
// 				return
// 			}
// 		case <-s.quitch:
// 			return
// 		}
// 	}
// }

func main() {
	server := NewServer(":8989")
	go func() {
		fmt.Println("5")
		for msg := range server.msgch {
			fmt.Println("5.5")

			fmt.Print(string(msg))
		}
	}()

	log.Fatal(server.Start())
}
