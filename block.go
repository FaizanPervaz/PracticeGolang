package main

import (
	"fmt"
	"time"
)

type Block struct {
	Index     int
	Timestamp time.Time
	Data      string
	PrevHash  string
	Hash      string
}

type Blockchain struct {
	Blocks []Block
}

func DisplayAllBlocks(chain Blockchain) {
	for _, block := range chain.Blocks {
		fmt.Printf("Block #%d\n", block.Index)
		fmt.Printf("Timestamp: %s\n", block.Timestamp.Format(time.RFC3339))
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("PrevHash: %s\n", block.PrevHash)
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Println("--------------------------------------")
	}
}

func NewBlock(data string, prevHash string, chain Blockchain) Block {
	// Generate a new block with the given data and previous block hash
	newIndex := len(chain.Blocks) + 1
	newBlock := Block{
		Index:     newIndex,
		Timestamp: time.Now(),
		Data:      data,
		PrevHash:  prevHash,
		Hash:      "", // Placeholder for hash calculation
	}
	newBlock.Hash = calculateHash(newBlock)
	return newBlock
}

func ModifyBlock(block Block, newData string) Block {
	// Modify the data of a block
	block.Data = newData
	// Recalculate the hash
	block.Hash = calculateHash(block)
	return block
}

func calculateHash(block Block) string {
	// This is a placeholder function for calculating the hash of a block
	// In a real-world scenario, you'd use a cryptographic hash function like SHA256
	return "fakehash"
}

func main() {
	// Main function
}
