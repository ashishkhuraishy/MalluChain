package blockchain

import (
	"log"

	"github.com/boltdb/bolt"
)

const dbName = "database/blochain.db"
const dbBucket = "malluchain"

// BlockChain data structure
type BlockChain struct {
	Tip []byte
	Db  *bolt.DB
}

// AddBlock functions helps us to add a new data into the blockChain
func (bc *BlockChain) AddBlock(data string) {
	var lastHash []byte

	err := bc.Db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(dbBucket))
		lastHash = b.Get([]byte{'l'})

		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	newBlock := NewBlock(data, lastHash)

	bc.Tip = newBlock.Hash
	err = bc.Db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(dbBucket))
		b.Put([]byte{'l'}, bc.Tip)
		b.Put(newBlock.Hash, newBlock.Serialize())

		return nil
	})

	if err != nil {
		log.Panic(err)
	}

}

// NewBlockChain return a new blockchain for start with
func NewBlockChain() *BlockChain {
	var tip []byte

	db, err := bolt.Open(dbName, 0600, nil)

	if err != nil {
		log.Panic(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(dbBucket))

		if b == nil {
			b, _ = tx.CreateBucket([]byte(dbBucket))
			block := GenusisBlock()
			err = b.Put(block.Hash, block.Serialize())
			err = b.Put([]byte{'l'}, block.Hash)
			tip = block.Hash
		} else {
			tip = b.Get([]byte{'l'})
		}

		if err != nil {
			log.Panic(err)
		}

		return nil
	})

	return &BlockChain{tip, db}
}

// Iterator used to iterate
// through the db
type Iterator struct {
	currentHash []byte
	db          *bolt.DB
}

// NewIterator ...
func (bc *BlockChain) NewIterator() *Iterator {
	bci := Iterator{bc.Tip, bc.Db}

	return &bci
}

// Next returns the next block in the chain
func (i *Iterator) Next() *Block {
	var block *Block

	err := i.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(dbBucket))
		encodedBlock := b.Get(i.currentHash)
		block = Deserialize(encodedBlock)

		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	i.currentHash = block.PrevHash

	return block
}
