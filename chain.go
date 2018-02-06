package main

// Block represents a collection of
// transactions have have been verified
// by a validator/miner
type Block struct {
	Txs        []Tx
	MerkleRoot string
}

// Chain represents a list of blocks
// that have been accepted by the
// network
type Chain struct {
	Blocks []Block `json:"blocks"`
	Height int     `json:"height"`
}

// NewChain initialises a Chain
func NewChain() Chain {
	return Chain{}
}
