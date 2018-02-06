package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	state State
	chain Chain
)

// GetState handles requests to the /state
// route that displays the networks shared
// state machine object
func GetState(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(State{})
}

// GetChain handles requests to the /chain
// route that displays the current chain
// held by the serving node
func GetChain(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(chain)
}

func main() {
	// Initialise state with Genesis
	state = State{
		[]Account{
			Account{PublicKey: "", Address: "", Balance: 10},
		},
	}

	// Initialise chain
	chain = NewChain()

	// Initialise HTTP Server
	router := mux.NewRouter()
	router.HandleFunc("/state", GetState)
	router.HandleFunc("/chain", GetChain)

	log.Fatal(http.ListenAndServe(":8080", router))
}
