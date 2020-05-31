package blockchain

import (
	"bytes"
	"encoding/gob"
	"log"
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

// GenusisBlock the fisrt block in the Block chain
func GenusisBlock() *Block {
	return NewBlock("Genusis block", []byte{})
}

// Serialize the blocks to byte format
func (b *Block) Serialize() []byte {
	var result bytes.Buffer

	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(b)

	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}

// Deserialize a block from bytes
func Deserialize(b []byte) *Block {
	var block *Block

	decoder := gob.NewDecoder(bytes.NewReader(b))
	err := decoder.Decode(&block)

	if err != nil {
		log.Panic(err)
	}

	return block
}
