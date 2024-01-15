package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

// Transaction represents a blockchain transaction
type Transaction struct {
	ID        string    `json:"id"`
	Timestamp time.Time `json:"timestamp"`
	Sender    string    `json:"sender"`
	Recipient string    `json:"recipient"`
	Amount    float64   `json:"amount"`
	Signature string    `json:"signature"` // Placeholder for digital signature
}

// calculateHash generates a unique hash for the transaction
func (t *Transaction) calculateHash() {
	data := t.Sender + t.Recipient + string(t.Amount) + t.Timestamp.String()
	hash := sha256.Sum256([]byte(data))
	t.ID = hex.EncodeToString(hash[:])
}

// NewTransaction creates a new Transaction
func NewTransaction(sender, recipient string, amount float64) *Transaction {
	transaction := &Transaction{
		Sender:    sender,
		Recipient: recipient,
		Amount:    amount,
		Timestamp: time.Now(),
	}
	transaction.calculateHash() // Calculate the transaction's hash
	return transaction
}
