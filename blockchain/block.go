package blockchain

import (
	"bytes"
	"crypto/sha256"
	"strconv"
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
}

// SetHash is used to create a hash for the current block
func (b *Block) SetHash() {
	timeStamp := []byte(strconv.FormatInt(b.TimeStamp, 10))
	headers := bytes.Join([][]byte{b.PrevHash, b.Data, timeStamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}

// NewBlock creates a new block with the given data and previous hash
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), prevBlockHash, []byte(data), []byte{}}
	block.SetHash()

	return block
}
