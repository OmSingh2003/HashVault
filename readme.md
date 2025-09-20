# HashVault

**HashVault** is a decentralized and fully distributed content-addressable file storage system built in Go.

---

##  Current Status

**Under active development.**  
The foundational peer-to-peer TCP transport layer is currently being implemented.

---

## ✨ Features

### ✅ Current
- **Decentralized Storage**  
  No central server; files are stored across a network of peers.
- **Content-Addressable**  
  Files are identified and retrieved by the hash of their content, ensuring data integrity and de-duplication.
- **Peer-to-Peer Communication**  
  Utilizes a custom TCP-based transport layer for direct communication between nodes.

### 🔜 Planned
- **File Encryption**  
  Secure storage of files using encryption.
- **File Replication & Sharding**  
  Distributes and replicates file segments across the network for fault tolerance.
- **Streaming**  
  Support for streaming data to and from the network.

---

## 🔧 Prerequisites

- Go (latest stable version recommended, e.g., `1.21+`)

---

## 🛠️ Building the Project

To compile the project and generate the `fs` executable in the `bin/` directory:

```bash
make build
```

⸻

🚀 Running the Project

To start the application (currently launches the TCP transport listener):
``` bash
make run
```
This will execute ./bin/fs. Functionality will expand as development progresses.

⸻

🧪 Running Tests

To run all unit tests with verbose output:
```bash
make test
```

⸻

📁 Project Structure

Path	Description
main.go	Main entry point of the application.
p2p/	Peer-to-peer networking logic, including TCP transport.
Makefile	Automates build, run, and test tasks.
go.mod	Go module definition.
go.sum	Checksums for Go module dependencies.
bin/fs	Compiled executable (after make build).


⸻

📦 Dependencies
	•	Go Standard Library
	•	stretchr/testify
Used for unit test assertions (assert package).


