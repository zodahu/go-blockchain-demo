package main

import (
	"fmt"
	"publicChain/block"
)

func main() {
	bc := block.CreateGenesisBlockchain()
	fmt.Println(bc.Blocks[0])

	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	bc.AddBlockToBlockchain(prevBlock.Height+1, prevBlock.Hash, "Send $100 to B")

	prevBlock = bc.Blocks[len(bc.Blocks)-1]
	bc.AddBlockToBlockchain(prevBlock.Height+1, prevBlock.Hash, "Send $100 to C")

	prevBlock = bc.Blocks[len(bc.Blocks)-1]
	bc.AddBlockToBlockchain(prevBlock.Height+1, prevBlock.Hash, "Send $100 to D")

	blockBytes := bc.Blocks[0].Serialize()
	fmt.Println(blockBytes)

	b := block.Deserialize(blockBytes)
	fmt.Println(b)

}
