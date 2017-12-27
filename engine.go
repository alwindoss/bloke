package main

import (
	"log"
	"net/http"
)

// StartBloke starts an HTTP API Server to accept blockchain requests
func StartBloke(port string) {
	router := NewRouter()
	block, err := CreateGenesisBlock()
	if err != nil {
		log.Fatalf("unable to create the genesis block")
	}
	index = 0
	blockChain = append(blockChain, block)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
