// HashVault/p2p/tcp_transport.go
package p2p

import (
	"fmt"
	"net"
	"sync"
)

// TCPeer represents the remote node over a TCP established connection
type TCPeer struct {
	// conn is the underlying connection of the peer
	conn net.Conn
	// if we dial and retrieve a conn
	outbound bool
}

// newTCPeer creates and returns a new TCPeer instance.
func newTCPeer(conn net.Conn, outbound bool) *TCPeer {
	return &TCPeer{
		conn:     conn,
		outbound: outbound,
	}
}

// TCPTransport holds the state for the TCP transport layer.
type TCPTransport struct {
	ListenAddress string
	Listener      net.Listener
	mu            sync.RWMutex // To protect shared access to peers map
	peers         map[net.Addr]Peer
}

// NewTCPTransport creates a new TCPTransport instance.
func NewTCPTransport(listenerAddr string) *TCPTransport {
	return &TCPTransport{
		ListenAddress: listenerAddr,
		peers:         make(map[net.Addr]Peer), // Initialize the peers map
	}
}

// ListenAndAccept sets up the listener and starts the loop for accepting connections.
func (t *TCPTransport) ListenAndAccept() error {
	var err error
	t.Listener, err = net.Listen("tcp", t.ListenAddress)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", t.ListenAddress, err)
	}

	fmt.Printf("TCP transport listening on %s\n", t.ListenAddress)

	// Start the accept loop in a new goroutine so ListenAndAccept can return.
	go t.startAcceptLoop()

	return nil
}

// startAcceptLoop runs a loop that accepts incoming connections.
func (t *TCPTransport) startAcceptLoop() {
	// It's good practice to ensure the listener is closed if the accept loop stops for any reason.
	// defer t.Listener.Close() // Consider the best place for this based on overall application lifecycle.

	for {
		conn, err := t.Listener.Accept()
		if err != nil {
			// Handle common errors, like when the listener is closed.
			if opError, ok := err.(*net.OpError); ok && opError.Err.Error() == "use of closed network connection" {
				fmt.Println("Listener closed, accept loop stopping.")
				return // Exit loop if listener is closed
			}
			fmt.Printf("TCP accept error: %s\n", err)
			continue // Continue to try accepting other connections
		}
		// When a connection is accepted, handle it in a new goroutine.
		go t.handleConn(conn)
	}
}

// handleConn is called for each accepted connection.
func (t *TCPTransport) handleConn(conn net.Conn) {
	// You can create a peer object if you need to manage state for this connection.
	// Since the connection is accepted by our listener, 'outbound' is false.
	_ = newTCPeer(conn, false) // The peer variable isn't used further in this specific example,
	                            // but creating it might be part of your broader design.

	// Print the desired message to the console.
	fmt.Println("1 incomming connection") // Note: "incomming" is as per your request.

	// At this point, the connection `conn` is open.
	// If you want to immediately close it after printing, you would add:
	// conn.Close()
	// However, for `telnet` to stay connected until you manually close it from the telnet client,
	// you should leave `conn.Close()` out of here or handle the connection lifecycle more explicitly
	// (e.g., reading data from it in a loop, and closing when done or on error).
}

// Interface definitions (ensure these are in your p2p/transport.go or accessible)
// type Peer interface{}
// type Transport interface{
// ListenAndAccept() error
// }
