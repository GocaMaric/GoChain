package api

import (
	"encoding/json"
	"github.com/GocaMaric/GoChain/pkg/blockchain"
	"github.com/gorilla/mux"
	"net/http"
)

// APIServer represents the REST API server
type APIServer struct {
	Blockchain *blockchain.Blockchain
}

// NewAPIServer creates a new API server
func NewAPIServer(bc *blockchain.Blockchain) *APIServer {
	return &APIServer{Blockchain: bc}
}

// Start starts the API server
func (api *APIServer) Start(port string) {
	router := mux.NewRouter()
	router.HandleFunc("/blocks", api.getBlocks).Methods("GET")
	http.ListenAndServe(":"+port, router)
}

// getBlocks handles the blocks retrieval API request
func (api *APIServer) getBlocks(w http.ResponseWriter, r *http.Request) {
	blocks := api.Blockchain.Blocks
	json.NewEncoder(w).Encode(blocks)
}
