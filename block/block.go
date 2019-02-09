package block

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
	"strconv"
	"time"
)

// Block definition
type Block struct {
	Height        int64
	PrevBlockHash []byte
	Data          []byte
	Timestamp     int64
	Hash          []byte
	Nonce         int64
}

// CreateGenesisBlock creates the genesis block
func CreateGenesisBlock(data string) *Block {
	prevBlockHash := make([]byte, 32)
	return CreateBlock(1, prevBlockHash, data)
}

// CreateBlock creates a new block
func CreateBlock(height int64, prevBlockHash []byte, data string) *Block {
	b := &Block{Height: height, PrevBlockHash: prevBlockHash, Data: []byte(data), Timestamp: time.Now().Unix(), Hash: nil, Nonce: 0}

	pow := NewProofOfWork(b)
	hash, nonce := pow.Run()

	b.Hash = hash[:]
	b.Nonce = nonce

	return b
}

// Serialize converts block to []byte
func (b *Block) Serialize() []byte {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)

	err := encoder.Encode(b)
	if err != nil {
		log.Panic(err)
	}

	return res.Bytes()
}

// Deserialize converts []byte to block
func Deserialize(blockBytes []byte) *Block {
	var b Block

	decoder := gob.NewDecoder(bytes.NewReader(blockBytes))
	err := decoder.Decode(&b)
	if err != nil {
		log.Panic(err)
	}

	return &b
}

// SetHash is unused
func (b *Block) SetHash() {
	// Covert elements to []byte

	hBytes := Int64ToHex(b.Height)

	tsString := strconv.FormatInt(b.Timestamp, 2)
	tsBytes := []byte(tsString)

	blockBytes := bytes.Join([][]byte{hBytes, b.PrevBlockHash, b.Data, tsBytes, b.Hash}, []byte{})

	hash := sha256.Sum256(blockBytes)
	b.Hash = hash[:]

}
