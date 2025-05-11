package p2p // Declares the test file is part of the p2p package.

import (
	// Imports the standard Go networking library for functions like Listen, Dial, Addr types.
	"testing" // Imports the standard Go testing library, providing *testing.T for test context.

	// Imports the testify/assert package, a popular third-party library for writing assertions in tests.
	"github.com/stretchr/testify/assert"
)

// TestTCPTransport defines a test function for the TCPTransport type.
func TestTCPTransport(t *testing.T) {
	// Define the network address string that the TCPTransport instance is intended to use for listening.
	listAddr := ":4000"

	// Create a new TCPTransport instance using the constructor function NewTCPTransport.
	// The 'tr' variable holds a pointer to this instance, which represents the TCP transport layer being tested.
	tr := NewTCPTransport(listAddr)

	// === Assertions on the initial state of the 'tr' (TCPTransport) instance ===

	// Assert that the ListenAddress field of the created 'tr' instance matches the 'listAddr' provided to the constructor.
	// Verifies: The constructor correctly stored the intended listen address.
	assert.Equal(t, tr.ListenAddress, listAddr)

//Server
	tr.ListenAndAccept()
}

