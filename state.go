package main

// Account stores the details of
// a single account on the coil
// network
type Account struct {
	Address   string  `json:"address"`
	PublicKey string  `json:"publickey"`
	Balance   float64 `json:"balance"`
}

// State stores the state of our
// application
type State struct {
	Accounts []Account `json:"accounts"`
}
