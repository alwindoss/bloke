package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type payload struct {
	Index string `json:"index"`
	Data  string `json:"data"`
}

func addBlockChanin(w http.ResponseWriter, r *http.Request) {
	var p payload
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &p); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	prevBlock := blockChain[len(blockChain)-1]
	prevHash := prevBlock.Hash
	block, err := NewBlock(p.Index, time.Now().String(), p.Data, prevHash)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(500) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	blockChain = append(blockChain, &block)
}

func getBlockchain(w http.ResponseWriter, r *http.Request) {
	blockByte, err := json.Marshal(blockChain)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	w.Write(blockByte)
}
