package main

import (
	"demochain/core"
	"encoding/json"
	"io"
	"net/http"
)

var blockchain *core.Blockchain

func run() {
	http.HandleFunc("/blockchain/get", blockchainGetHandler)
	http.HandleFunc("/blockchain/write", blockchainWriterHandler)
	http.ListenAndServe("localhost:8888", nil)
}

func blockchainGetHandler(w http.ResponseWriter, r *http.Request) {
	bytes, error := json.Marshal(blockchain)
	if error != nil {
		http.Error(w, error.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(bytes))
}

func blockchainWriterHandler(w http.ResponseWriter, r *http.Request) {
	blockData := r.URL.Query().Get("data")
	blockchain.SendDdata(blockData)
	blockchainGetHandler(w, r)
}

func main() {
	blockchain = core.NewBlockchain()
	run()
}
