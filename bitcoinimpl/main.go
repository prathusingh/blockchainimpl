package main

import (
	"time"
	"fmt"
)

// Block structure
type Block struct {
	Index int // Block number
	Timestamp time.Time // Block creation time
	Data string // Data (Eg transaction details)
	PrevHash string // Hash of the previous block
	Hash string // Current block hash
	Nonce int // Proof of work nonce
}

// Blockchain structure
type Blockchain struct {
	Chain []Block
	CurrentTransactions []string
}






func main() {
	// Create the Genesis Block
	genesisBlock := NewBlock([]string{"Genesis Block"}, "0")
	fmt.Println(genesisBlock)
}




