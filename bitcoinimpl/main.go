package main

import (
	"time"
	"fmt"
	"crypto/sha256"
	"encoding"
)

// Block structure
type Block struct {
	Index int // Block number
	Timestamp string // Block creation time
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

// Generates the has og the block
func (b *Block) CalculateHash() string {
	record := fmt.Sprintf("%d%s%s%s%d", b.Index, b.Timestamp, b.Data, b.PrevHash, b.Nonce)
	hash := sha256.New()
	hash.Write([]byte(record))
	return fmt.Sprintf("%x", hash.Sum(nil))
}
 
// New Block creates a new block with proof of work
func NewBlock(prevBlock Block, data string) Block {
	block := Block{
		Index: prevBlock.Index + 1,
		Timestamp: time.Now().String(),
		Data: data,
		PrevHash: prevBlock.Hash,
		Nonce: 0,
	}

	block.Hash = ""
	return block
}

func CreateGenesisBlock() Block {
	return Block{
		Index: 0,
		Timestamp: time.Now().String(),
		Data: "Genesis Block",
		PrevHash: "",
		Hash: "",
		Nonce: 0,
	}
}


// Initiates the blockchain with the genesis block
func NewBlockchain() *Blockchain {
	genesisBlock := CreateGenesisBlock()
	return &Blockchain{
		Chain: []Block{genesisBlock},
	}
}

func (bc *Blockchain) AddBlock(data string) {
	prev
}






func main() {
	// Create the Genesis Block
	genesisBlock := NewBlock([]string{"Genesis Block"}, "0")
	fmt.Println(genesisBlock)
}




