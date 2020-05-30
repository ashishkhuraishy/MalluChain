package main

import (
	"fmt"

	"github.com/ashishkhuraishy/mallu-chain/blockchain"
)

func main() {
	fmt.Println("Mallu Chain starting...")

	bc := blockchain.NewBlockChain()
	bc.AddBlock("First Node")
	bc.AddBlock("Second Node")

	for _, block := range bc.Blocks {
		fmt.Printf("Hash : %x\n", block.PrevHash)
		fmt.Printf("Data : %s\n", block.Data)
		fmt.Printf("Hash : %x\n\n", block.Hash)
	}

}
