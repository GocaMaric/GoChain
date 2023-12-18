package network

import (
	"log"
	"net"
)

// Node represents a node in the blockchain network
type Node struct {
	IP   string
	Port string
}

// Start initializes the network node
func (node *Node) Start() {
	listener, err := net.Listen("tcp", node.IP+":"+node.Port)
	if err != nil {
		log.Fatalf("Unable to start node: %v", err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Unable to accept connection: %v", err)
			continue
		}

		// Handle connection
		go handleConnection(conn)
	}
}

// handleConnection handles incoming network connections
func handleConnection(conn net.Conn) {
	// Network handling logic
	conn.Close()
}
