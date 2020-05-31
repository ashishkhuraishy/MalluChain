package main

import (
	"fmt"
	"strconv"

	"github.com/ashishkhuraishy/mallu-chain/blockchain"
)

func main() {
	fmt.Println("Mallu Chain starting...")

	bc := blockchain.NewBlockChain()
	bc.AddBlock("First Node")
	bc.AddBlock("Second Node")

	bci := bc.NewIterator()

	for {

		block := bci.Next()

		fmt.Printf("Previous Hash : %x\n", block.PrevHash)
		fmt.Printf("Data : %s\n", block.Data)
		fmt.Printf("Nonce : %d\n", block.Nonce)
		pow := blockchain.NewProofOfWork(block)
		fmt.Printf("Validate : %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Printf("Hash : %x\n\n", block.Hash)

		if len(block.PrevHash) == 0 {
			break
		}
	}

}
