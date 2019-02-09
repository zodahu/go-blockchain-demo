package block

import (
	"bytes"
	"crypto/sha256"
	"math/big"
)

const targetBit = 16

// ProofOfWork .
type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

// NewProofOfWork .
func NewProofOfWork(block *Block) *ProofOfWork {
	target := big.NewInt(1)
	target = target.Lsh(target, 256-targetBit)
	return &ProofOfWork{block, target}
}

// Run .
func (p *ProofOfWork) Run() ([]byte, int64) {
	var nonce int64
	var hashInt big.Int
	var hash [32]byte
	for {
		bytesData := p.blockToBytes(nonce)
		hash = sha256.Sum256(bytesData)
		hashInt.SetBytes(hash[:])

		if p.Target.Cmp(&hashInt) == 1 {
			break
		}
		nonce++
	}

	return hash[:], nonce
}

// IsValid .
func (p *ProofOfWork) IsValid() bool {
	var hashInt big.Int
	hashInt.SetBytes(p.Block.Hash)

	if p.Target.Cmp(&hashInt) == 1 {
		return true
	}

	return false
}

func (p *ProofOfWork) blockToBytes(nonce int64) []byte {
	bytesData := bytes.Join(
		[][]byte{
			Int64ToHex(p.Block.Height),
			p.Block.Hash,
			p.Block.Data,
			Int64ToHex(p.Block.Timestamp),
			Int64ToHex(nonce),
		},
		[]byte{},
	)

	return bytesData
}
