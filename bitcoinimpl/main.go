package main

import (
	"time"
	"fmt"
	"crypto/sha256"
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

// Adds a new block to the blockchain
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.Chain[len(bc.Chain)-1]
	newBlock := NewBlock(prevBlock,data)
	bc.Chain = append(bc.Chain, newBlock)
}

// checks if the hash is valid based on difficulty (leading zeroes)
func isValidHash(hash string, difficulty int) bool {
	prefix := string(make([]byte, difficulty))
	return hash[:difficulty] == prefix
}

// Tries different nonces to find a valid hash
func ProofOfWork(b *Block, difficulty int) {
	for {
		hash := b.CalculateHash()
		if isValidHash(hash, difficulty) {
			break
		}
		b.Nonce++
	}
}


func main() {
	blockchain := NewBlockchain()

	difficulty := 2 // difficulty level for Proof of Work

	for i := 1; i <=3; i++ {
		blockchain.AddBlock(fmt.Sprintf("Transaction %d", i))

		// Get the last block
		lastBlock := blockchain.Chain[len(blockchain.Chain)-1]

		// Perform Proof Of Work on the new block
		ProofOfWork(&lastBlock, difficulty)

		fmt.Printf("Block #%d\n", lastBlock.Index)
		fmt.Printf("Timestamp: %s\n", lastBlock.Timestamp)
		fmt.Printf("Data: %s\n", lastBlock.Data)
		fmt.Printf("PrevHash: %s\n", lastBlock.PrevHash)
		fmt.Printf("Hash: %s\n", lastBlock.Hash)
		fmt.Printf("Nonce: %d\n\n", lastBlock.Nonce)
	}
}




