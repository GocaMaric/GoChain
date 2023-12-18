// transaction.go
package blockchain

// Transaction represents a blockchain transaction
type Transaction struct {
	Sender    string
	Recipient string
	Amount    float64
}

// NewTransaction creates a new Transaction
func NewTransaction(sender, recipient string, amount float64) *Transaction {
	return &Transaction{Sender: sender, Recipient: recipient, Amount: amount}
}
