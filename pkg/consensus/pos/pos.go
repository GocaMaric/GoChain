package pos

import (
    "github.com/GocaMaric/GoChain/pkg/blockchain"
    "math/rand"
    "time"
    "errors"
)

// Stakeholder represents a node in the network with its stake value
type Stakeholder struct {
    NodeID string
    Stake  int
}

// ProofOfStake represents the Proof of Stake consensus mechanism
type ProofOfStake struct {
    blockchain  *blockchain.Blockchain
    stakeholders []Stakeholder
}

// NewProofOfStake creates a new ProofOfStake instance
func NewProofOfStake(bc *blockchain.Blockchain) *ProofOfStake {
    return &ProofOfStake{
        blockchain:   bc,
        stakeholders: make([]Stakeholder, 0),
    }
}

// ValidateBlock validates a block based on the Proof of Stake mechanism
func (pos *ProofOfStake) ValidateBlock(block *blockchain.Block) bool {
    // Placeholder for PoS block validation logic
    // In a real implementation, this would involve checking the stake of the block creator
    return true
}

// SelectBlockCreator selects a node as the block creator based on their stake
func (pos *ProofOfStake) SelectBlockCreator() string {
    totalStake := 0
    for _, stakeholder := range pos.stakeholders {
        totalStake += stakeholder.Stake
    }

    if totalStake == 0 {
        return "" // No stakeholders to select from
    }

    rand.Seed(time.Now().UnixNano())
    winningPoint := rand.Intn(totalStake)
    currentPoint := 0

    for _, stakeholder := range pos.stakeholders {
        currentPoint += stakeholder.Stake
        if currentPoint >= winningPoint {
            return stakeholder.NodeID
        }
    }

    return ""
}

// CreateNewBlock creates a new block to be added to the blockchain
func (pos *ProofOfStake) CreateNewBlock() (*blockchain.Block, error) {
    nodeID := pos.SelectBlockCreator()
    if nodeID == "" {
        return nil, errors.New("no block creator selected")
    }

    // Create a new block. In a real implementation, you would handle transactions,
    // state updates, and other blockchain-specific logic here.
    newBlock := blockchain.NewBlock("New PoS Block", []byte{}, pos.blockchain.GetLatestBlock().Hash)
    return newBlock, nil
}

// UpdateStakeholder updates the stake for a given node
func (pos *ProofOfStake) UpdateStakeholder(nodeID string, stake int) {
    for i, stakeholder := range pos.stakeholders {
        if stakeholder.NodeID == nodeID {
            pos.stakeholders[i].Stake = stake
            return
        }
    }
    // If the stakeholder doesn't exist, add them to the list
    pos.stakeholders = append(pos.stakeholders, Stakeholder{NodeID: nodeID, Stake: stake})
}
