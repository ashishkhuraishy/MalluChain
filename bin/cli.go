package bin

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/ashishkhuraishy/mallu-chain/blockchain"
)

// CLI used to run args using cmd
type CLI struct {
	Bc *blockchain.BlockChain
}

func (cli *CLI) printUsage() {
	fmt.Println("Usage")
	fmt.Println(" addblock -data (Adds a new block to the chain)")
	fmt.Println(" printblocks (Displays all the block in the chain)")

}

func (cli *CLI) validate() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}

func (cli *CLI) addBlock(data string) {
	cli.Bc.AddBlock(data)
	fmt.Println("Sucess")
}

func (cli *CLI) printChain() {
	bci := cli.Bc.NewIterator()

	for {
		block := bci.Next()

		fmt.Printf("Prev Hash : %x\n", block.PrevHash)
		fmt.Printf("Data : %s\n", block.Data)
		fmt.Printf("Hash : %x\n\n", block.Hash)

		if len(block.PrevHash) == 0 {
			break
		}
	}
}

// Run the Cli
func (cli *CLI) Run() {
	cli.validate()

	addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
	printblocks := flag.NewFlagSet("printblocks", flag.ExitOnError)

	addBlockData := addBlockCmd.String("data", "", "Add this data to a new block")

	switch os.Args[1] {
	case "addblock":
		err := addBlockCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "printblocks":
		err := printblocks.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		cli.printUsage()
		os.Exit(1)
	}

	if addBlockCmd.Parsed() {
		if *addBlockData == "" {
			cli.printUsage()
			os.Exit(1)
		}
		cli.addBlock(*addBlockData)
	}

	if printblocks.Parsed() {
		cli.printChain()
	}
}
