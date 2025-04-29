package p2p

import "net"

type TCPTransport struct {
	ListenAddress string
	Listener      net.Listener
	mu            syncRWMutex // Mutex protects peer
	peers         map[net.Addr]Peer
}

func NewTCPTransport(listener string) *TCPTransport {
	return &TCPTransport{ListenAdress: listener}
}

// func Tset()
// t := NewTCPTransport("localhost:8080")
