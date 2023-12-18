package blockchain

import (
	"encoding/hex"
	"log"
)

// Blockchain represents the chain of blocks
type Blockchain struct {
	Blocks []*Block
}

// AddBlock adds a new block to the blockchain
// Returns an error if it fails to create a new block
func (bc *Blockchain) AddBlock(data string) error {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	prevHash, err := hex.DecodeString(prevBlock.Hash)
	if err != nil {
		// Log the error and return it
		log.Printf("Failed to decode previous block hash: %v", err)
		return err
	}

	newBlock := NewBlock(data, prevHash)
	bc.Blocks = append(bc.Blocks, newBlock)
	return nil
}

// NewBlockchain creates a new blockchain with genesis Block
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewBlock("Genesis Block", []byte{})}}
}
