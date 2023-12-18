// block.go
package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"time"
)

// Block represents a single block in the blockchain
type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          string
}

// SetHash calculates and sets the hash of the block
func (b *Block) SetHash() {
	timestamp := []byte(time.Unix(b.Timestamp, 0).String())
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hex.EncodeToString(hash[:])
}

// NewBlock creates a new Block
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, ""}
	block.SetHash()
	return block
}
