// Chat state and broadcasting
package chat

import (
	"fmt"
	"net"
	"sync"

	"tcp-chat/client"
)

type ChatRoom struct {
	clients map[string]*client.Client
	history []string
	mu      sync.Mutex
}

func (c *ChatRoom) BroadcastMessage(sender, msg string) {
	timestamp := GetTimestamp()
	full := fmt.Sprintf("%s[%s]: %s", timestamp, sender, msg)
	c.mu.Lock()
	defer c.mu.Unlock()
	for _, cli := range c.clients {
		if cli.Name != sender {
			fmt.Fprintln(cli.Conn, full)
		}
	}

	c.history = append(c.history, full)
}

func Register(name string) {
	// TODO: Add to client map
}

func RemoveClient(name string) {
	// TODO: Remove from client map and broadcast leave message
}

func (c *ChatRoom) SendHistory(conn net.Conn) {
	for _, msg := range c.history {
		fmt.Fprintln(conn, msg)
	}
}
