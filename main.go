// Entry point
package main

import (
	"log"
	"os"

	"tcp-chat/server"
)

func main() {
	if len(os.Args) > 2 {
		log.Fatal("Usage: go run main.go [address]")
	}
	addr := ":8989"
	if len(os.Args) == 2 {
		addr = ":" + os.Args[1]
	}

	s := server.NewServer(addr)
	log.Fatal(s.Start())
}
