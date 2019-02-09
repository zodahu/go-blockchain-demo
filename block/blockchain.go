package block

// Blockchain .
type Blockchain struct {
	Blocks []*Block
}

// CreateGenesisBlockchain .
func CreateGenesisBlockchain() *Blockchain {
	genesisBlock := CreateGenesisBlock("Genesis Block")

	return &Blockchain{[]*Block{genesisBlock}}
}

// AddBlockToBlockchain .
func (bc *Blockchain) AddBlockToBlockchain(height int64, prevBlockHash []byte, data string) {
	b := CreateBlock(height, prevBlockHash, data)
	bc.Blocks = append(bc.Blocks, b)
}
