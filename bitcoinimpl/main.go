package main

import (
	"time"
	"fmt"
)

// Block structure
type Block struct {
		Index int // Block number
		Timestamp time.Time // Block creation time
		Transactions []string // List of transactions
		PreviousHash string // Hash of the previous block
		Hash string // Current block hash
		Nonce int // Proof of work nonce
}

// Blockchain structure
var Blockchain []Block

// Create a new block
func NewBlock(transactions []string, previousHash string) Block {
		block := Block{
			Index: len(Blockchain),
			Timestamp: time.Now(),
			Transactions: transactions,
			PreviousHash: previousHash,
		}

		block.Hash, block.Nonce = "", 1

		return block
}


func main() {
	// Create the Genesis Block
	genesisBlock := NewBlock([]string{"Genesis Block"}, "0")
	fmt.Println(genesisBlock)
}




