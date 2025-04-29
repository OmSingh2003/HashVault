package p2p

import (
	"net"
	"sync"
)

type TCPTransport struct {
	ListenAddress string
	Listener      net.Listener
	mu            sync.RWMutex // Mutex protects peer
	peers         map[net.Addr]Peer
}

func NewTCPTransport(listener string) *TCPTransport {
	return &TCPTransport{ListenAddress: listener}
}

// func Tset()
// t := NewTCPTransport("localhost:8080")
