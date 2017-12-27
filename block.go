package main

import (
	"crypto/sha256"
	"fmt"
	"log"
	"strconv"
	"time"
)

var blockChain []*Block
var index int

// Block holds the datastructure of the blockchain
type Block struct {
	Index        string `json:"index"`
	Timestamp    string `json:"time_stamp"`
	Data         string `json:"data"`
	PreviousHash string `json:"previous_hash"`
	Hash         string `json:"hash"`
}

// NewBlock is the factory to create a new block
func NewBlock(index, timestamp, data, previousHash string) (Block, error) {
	hasher := sha256.New()
	hasher.Write([]byte(index))
	hasher.Write([]byte(timestamp))
	hasher.Write([]byte(data))
	hasher.Write([]byte(previousHash))
	h := hasher.Sum(nil)
	hStr := fmt.Sprintf("%x", h)
	fmt.Println(fmt.Sprintf("%x", h))
	return Block{
		Index:        index,
		Timestamp:    timestamp,
		Data:         data,
		PreviousHash: previousHash,
		Hash:         hStr,
	}, nil
}

// CreateGenesisBlock creates the first block in the blockchain
func CreateGenesisBlock() (*Block, error) {
	block, err := NewBlock("0", time.Now().String(), "Genesis Block", "0")
	return &block, err
}

// NextBlock takes in the previous block and creates the next block in the blockchain
func NextBlock(prevBlock Block, data string) (*Block, error) {
	i, err := strconv.Atoi(prevBlock.Index)
	if err != nil {
		log.Printf("not a valid index: %v", err)
		return nil, err
	}
	newIndex := string(i + 1)
	block, err := NewBlock(newIndex, time.Now().String(), data, prevBlock.Hash)
	return &block, err
}
