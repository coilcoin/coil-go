package main

// Tx represents a single transaction
// of a given amount between two coil
// accounts
type Tx struct {
	Signature string `json:"signature"`
	PublicKey string `json:"publickey"`
	Amount    string `json:"amount"`
	From      string `json:"from"`
	To        string `json:"to"`
}
