package blockchain

import (
    "bytes"
    "crypto/sha256"
    "encoding/binary"
    "encoding/hex"
    "time"
)

// Block represents a single block in the blockchain
type Block struct {
    Timestamp     int64  `json:"timestamp"`
    Data          []byte `json:"data"`
    PrevBlockHash []byte `json:"prev_block_hash"`
    Hash          string `json:"hash"`
    Nonce         int    `json:"nonce"`
}

// SetHash calculates and sets the hash of the block
func (b *Block) SetHash() {
    var headers []byte
    headers = append(headers, b.PrevBlockHash...)
    headers = append(headers, b.Data...)
    headers = append(headers, intToHex(b.Timestamp)...)
    headers = append(headers, intToHex(int64(b.Nonce))...)

    hash := sha256.Sum256(headers)
    b.Hash = hex.EncodeToString(hash[:])
}

// intToHex converts an int64 to a byte array
func intToHex(num int64) []byte {
    buff := new(bytes.Buffer)
    err := binary.Write(buff, binary.BigEndian, num)
    if err != nil {
        // In a real application, you should handle this error appropriately
        panic("Failed to convert int to hex: " + err.Error())
    }
    return buff.Bytes()
}

// NewBlock creates a new Block
func NewBlock(data string, prevBlockHash []byte) *Block {
    block := &Block{
        Timestamp:     time.Now().Unix(),
        Data:          []byte(data),
        PrevBlockHash: prevBlockHash,
        Nonce:         0, // Nonce should be determined by a mining function in PoW
    }
    block.SetHash() // Hash is set based on the provided data
    return block
}
