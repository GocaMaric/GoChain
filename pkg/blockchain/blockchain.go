package blockchain

import (
    "encoding/hex"
    "log"
    "errors"
)

// Blockchain represents the chain of blocks
type Blockchain struct {
    Blocks []*Block
}

// AddBlock adds a new block to the blockchain
// Returns an error if it fails to create a new block
func (bc *Blockchain) AddBlock(data string) error {
    var prevHash []byte
    var err error

    // Handling the case when adding the first block after the genesis block
    if len(bc.Blocks) > 0 {
        prevBlock := bc.Blocks[len(bc.Blocks)-1]
        prevHash, err = hex.DecodeString(prevBlock.Hash)
        if err != nil {
            log.Printf("Failed to decode previous block hash: %v", err)
            return err
        }
    }

    newBlock := NewBlock(data, prevHash)
    bc.Blocks = append(bc.Blocks, newBlock)
    return nil
}

// NewBlockchain creates a new Blockchain with genesis Block
func NewBlockchain() *Blockchain {
    blockchain := &Blockchain{}
    genesisBlock := NewBlock("Genesis Block", []byte{})
    blockchain.Blocks = append(blockchain.Blocks, genesisBlock)
    return blockchain
}
