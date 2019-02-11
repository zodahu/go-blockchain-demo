package block

import (
	"bytes"
	"fmt"
	"log"
	"time"

	"github.com/boltdb/bolt"
)

const dbName = "go-blockchain-demo.db"
const blockTableName = "blocks"
const blockTableLast = "Last"

// Blockchain .
type Blockchain struct {
	Last []byte   // the latest block
	DB   *bolt.DB // each blockchain maps to one database
}

// CreateGetBlockchain will create genesis block if blockchain is empty,
// or get the latest blockchain
func CreateGetBlockchain() *Blockchain {
	var LastHash []byte

	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	db, err := bolt.Open(dbName, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {

		// try to find the table
		b := tx.Bucket([]byte(blockTableName))

		if b == nil { // create genesis block for empty DB
			b, err = tx.CreateBucket([]byte(blockTableName))
			if err != nil {
				log.Panic(err)
				return err
			}

			genesisBlock := CreateGenesisBlock("Genesis Block")
			LastHash = genesisBlock.Hash
			err = b.Put(LastHash, genesisBlock.Serialize())
			if err != nil {
				log.Panic(err)
				return err
			}

			err = b.Put([]byte(blockTableLast), LastHash)
			if err != nil {
				log.Panic(err)
				return err
			}
		} else { // assign the last block
			LastHash = b.Get([]byte(blockTableLast))
		}

		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	return &Blockchain{LastHash, db}
}

// AddBlockToBlockchain inserts block to DB
func (bc *Blockchain) AddBlockToBlockchain(data string) {

	err := bc.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockTableName))
		if b != nil {
			// retrieve the last block
			lastBlock := Deserialize(b.Get(bc.Last))

			bk := CreateBlock(lastBlock.Height+1, lastBlock.Hash, data)

			// Put new block to DB
			err := b.Put(bk.Hash, bk.Serialize())
			if err != nil {
				log.Panic(err)
				return err
			}

			// Update Last in blockchain and DB
			err = b.Put([]byte(blockTableLast), bk.Hash)
			if err != nil {
				log.Panic(err)
				return err
			}
			bc.Last = bk.Hash
		}
		return nil
	})

	if err != nil {
		log.Panic(err)
	}
}

// PrintBlockchain prints blocks
func (bc *Blockchain) PrintBlockchain() {
	var bk *Block
	currHash := bc.Last
	for {
		err := bc.DB.View(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte(blockTableName))
			if b != nil {
				bk = Deserialize(b.Get(currHash))
				fmt.Printf("%d\n", bk.Height)
				fmt.Printf("%x\n", bk.PrevBlockHash)
				fmt.Printf("%s\n", bk.Data)
				fmt.Printf("%s\n", time.Unix(bk.Timestamp, 0).Format("2006-01-02 03:04:05 PM"))
				fmt.Printf("%x\n", bk.Hash)
				fmt.Printf("%d\n", bk.Nonce)
				fmt.Println()
			}

			return nil
		})

		if err != nil {
			log.Panic(err)
			break
		}

		// check genesis block
		genesisPrevBlockHash := make([]byte, 32)
		if bytes.Equal(bk.PrevBlockHash, genesisPrevBlockHash) {
			break
		}

		// Retrieve the previous block
		currHash = bk.PrevBlockHash
	}
}
