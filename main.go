package main

import (
	"github.com/ashishkhuraishy/mallu-chain/bin"
	"github.com/ashishkhuraishy/mallu-chain/blockchain"
)

func main() {
	bc := blockchain.NewBlockChain()
	defer bc.Db.Close()

	cli := bin.CLI{bc}
	cli.Run()

}
