package p2p

// This file defines the TCPTransport struct and its methods.
import (
	"fmt"
	"net"  // a net package to handle network connections
	"sync" // a sync package to handle concurrency and avoid RACE conditons
)
// TCPeer represents the remote node over a TCP established connection
type TCPeer struct {
// conn is the undelying connection of the peer
conn net.Conn
// if we dial and retrieve a conn 
outbound bool
}
func newTCPeer(conn net.Conn , outbound bool) *TCPeer{
return &TCPeer {
conn :    conn
outbound : outbound
}
}
type TCPTransport struct { // Data structure for TCP tansport protocol
	ListenAddress string
	// Stores the network address on which transport instance
	// will listen for incoming peer connections
	Listener net.Listener
	// Holds the underlying TCP listener object returned by net.Listener
	//net.Listener is used for accepting incoming connections
	// It is an interface that provides a method to accept incoming connections
	// and returns a net.Conn object representing the connection
	// net.Conn is an interface that represents a connection to a network endpoint
	mu sync.RWMutex // Mutex protects peer
	//A write/Read mutex which is used to protect the peer
	//map from race conditons that could occur
	// if multiple goroutines try to read or write to the map simultaneously
	peers map[net.Addr]Peer
	//A map holding the currently connected peers.
}

func NewTCPTransport(listener string) *TCPTransport {
return &TCPTransport{ListenAddress: listener}
}
func (t *TCPTransport) ListenAndAccept() error {
	var err error
	t.Listener, err = net.Listen("tcp", t.ListenAddress)
	if err != nil {
		return err
	}
	return nil
}
func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.Listener.Accept()
		if err != nil {
			fmt.Printf("TCP accept error: %s\n", err) // fixed format string
			continue
		}
		go t.handleConn(conn) // optionally handle connection in a new goroutine
	}
}

// It would typically involve reading/writing data, handshaking, and managing peer state.
func (t *TCPTransport) handleConn(conn net.Conn) {
peer := NewTCPeer(conn , true)
	fmt.Printf(" New incoming connection: %s\n", peer
}

