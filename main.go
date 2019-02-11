package main

import (
	"publicChain/block"
)

func main() {

	bc := block.CreateGetBlockchain()
	defer bc.DB.Close()

	bc.AddBlockToBlockchain("new block")
	bc.PrintBlockchain()

}
