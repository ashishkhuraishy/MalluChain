package blockchain

import (
	"time"
)

// Block This struct is used to
// define the basic structure of a
// block in the chain
type Block struct {
	TimeStamp int64
	PrevHash  []byte
	Data      []byte
	Hash      []byte
	Nonce     int
}

// NewBlock creates a new block with the given data and previous hash
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), prevBlockHash, []byte(data), []byte{}, 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Nonce = nonce
	block.Hash = hash

	return block
}
