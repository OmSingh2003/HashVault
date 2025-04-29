package p2p

// Peer is an interface that represents the remote node.
type Peer interface{}

// Transport is an interface that defines the methods for a transport layer in a peer-to-peer network.
// It is responsible for establishing connections between peers and sending/receiving messages.
// Basically it handles the communication between the nodes in the network or you can say between peer -to -peer.
// This can be in form of TCP,UDP,WebSocket, etc.
type Transport interface{}
