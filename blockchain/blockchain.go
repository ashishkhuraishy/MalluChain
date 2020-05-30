package blockchain

// BlockChain data structure
type BlockChain struct {
	Blocks []*Block
}

// GenusisBlock the fisrt block in the Block chain
func GenusisBlock() *Block {
	return NewBlock("Genusis block", []byte{})
}

// AddBlock functions helps us to add a new data into the blockChain
func (bc *BlockChain) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)

	bc.Blocks = append(bc.Blocks, newBlock)
}

// NewBlockChain return a new blockchain for start with
func NewBlockChain() *BlockChain {
	return &BlockChain{[]*Block{GenusisBlock()}}
}
