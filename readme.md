```markdown
# HashVault

HashVault is a decentralized and fully distributed content-addressable file storage system built in Go.

## Current Status

Under active development. The foundational peer-to-peer TCP transport layer is currently being implemented.

## Features (Current and Planned)

* **Decentralized Storage:** No central server; files are stored across a network of peers.
* **Content-Addressable:** Files are identified and retrieved by the hash of their content, ensuring data integrity and de-duplication.
* **Peer-to-Peer Communication:** Utilizes a custom TCP-based transport layer for direct communication between nodes.
* **File Encryption (Planned):** Secure storage of files using encryption.
* **File Replication & Sharding (Planned):** Distributing and replicating file segments across the network for availability and fault tolerance.
* **Streaming (Planned):** Support for streaming data to and from the network.

## Prerequisites

* Go (latest stable version recommended, e.g., 1.21+)

## Building the Project

To build the project and generate the `fs` executable in the `bin/` directory:
```bash
make build
```

## Running the Project

To run the application (which currently starts the TCP transport listener):
```bash
make run
```
This will execute `./bin/fs`. The specific functionality of the running node will evolve as the project progresses.

## Running Tests

To execute the unit tests for the project:
```bash
make test
```
This command runs all tests in the project and provides verbose output.

## Project Structure

A brief overview of the key directories and files:

* `main.go`: The main entry point for the HashVault application.
* `p2p/`: Contains the peer-to-peer networking logic, including the TCP transport implementation.
* `Makefile`: Automates common tasks like building, running, and testing the project.
* `go.mod`, `go.sum`: Manage project dependencies.
* `bin/fs`: The compiled executable (after running `make build`).

## Dependencies

* Go Standard Library
* `github.com/stretchr/testify/assert`: Used for assertions in tests.
```
