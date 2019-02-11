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
	blc := &Block{Height: height, PrevBlockHash: prevBlockHash, Data: []byte(data), Timestamp: time.Now().Unix(), Hash: nil, Nonce: 0}

	pow := NewProofOfWork(blc)
	hash, nonce := pow.Run()

	blc.Hash = hash[:]
	blc.Nonce = nonce

	return blc
}

// Serialize converts block to []byte
func (blc *Block) Serialize() []byte {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)

	err := encoder.Encode(blc)
	if err != nil {
		log.Panic(err)
	}

	return res.Bytes()
}

// Deserialize converts []byte to block
func Deserialize(blockBytes []byte) *Block {
	var blc Block

	decoder := gob.NewDecoder(bytes.NewReader(blockBytes))
	err := decoder.Decode(&blc)
	if err != nil {
		log.Panic(err)
	}

	return &blc
}

// SetHash is unused currently
func (blc *Block) SetHash() {
	// Covert elements to []byte

	hBytes := Int64ToHex(blc.Height)

	tsString := strconv.FormatInt(blc.Timestamp, 2)
	tsBytes := []byte(tsString)

	blockBytes := bytes.Join([][]byte{hBytes, blc.PrevBlockHash, blc.Data, tsBytes, blc.Hash}, []byte{})

	hash := sha256.Sum256(blockBytes)
	blc.Hash = hash[:]

}
