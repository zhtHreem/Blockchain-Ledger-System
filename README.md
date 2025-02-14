 # Blockchain Ledger System

A simple command-line blockchain implementation in Go that demonstrates the core concepts of blockchain technology through a transaction ledger system.

## Project Overview

The Blockchain Ledger System allows users to create and manage a secure chain of transaction records. Each block is cryptographically linked, ensuring data integrity and immutability. The system uses SHA-256 hashing for security and supports basic consensus for validating the chain.
## Features

- **Blockchain Initialization**: Create a new blockchain with a genesis block
- **Transaction Management**: Add new transactions to the blockchain with sender, receiver, and amount details
- **Chain Integrity**: Verify the integrity of the entire blockchain using cryptographic hashes
- **Transaction History**: View all transactions in the blockchain in chronological order
- **Persistent Storage**: Data saved in JSON for continuity across sessions.
- **Simple CLI Interface**: Easy-to-use interface for interacting with the blockchain.


## Installation and Setup


1. **Clone the Repository**
   ```bash
   git clone https://github.com/zhtHreem/Blockchain-Ledger-System.git
   cd Blockchain-Ledger-System
   ```

2. **Build the Project**
   ```bash
   go build -o blockchain-ledger.exe cmd/ledger/main.go
   ```

3. **Basic Usage**
   ```bash
   # Initialize a new blockchain
   ./blockchain-ledger init

   # Add a new transaction
   ./blockchain-ledger add -from Alice -to Bob -amount 50.0

   # View all transactions
   ./blockchain-ledger list

   # Verify blockchain integrity
   ./blockchain-ledger verify
   ```

4. **Available Commands**
   ```bash
   init                          # Initialize a new blockchain
   add -from <addr> -to <addr> -amount <value>   # Add a new transaction
   verify                        # Verify the blockchain integrity
   list                         # List all blocks in the chain
   version                      # Show version information
   help                         # Show help message
   ```

## Tech Stack

- **Language**: Go 1.20+
- **Cryptography**: SHA-256 hashing (crypto/sha256)
- **Data Storage**: JSON file-based persistence
- **Architecture**: Command pattern for CLI operations
- **Build System**: Go modules for dependency management
