// HashVault/main.go
package main

import (
	"github.com/OmSingh2003/HashVault/p2p"
	"log"
)

func main() {
	tr := p2p.NewTCPTransport(":4000")
	if err := tr.ListenAndAccept(); err != nil {
		log.Fatalf("Failed to start transport: %v", err)
	}

	// Keep the main goroutine alive.
	// All connection handling happens in other goroutines.
	select {}

	// This line below select{} will not be reached.
	// fmt.Println("This will not print")
}
