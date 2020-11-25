package core

import (
	"github.com/boltdb/bolt"
	"log"
)

/**
让我们从 NewBlockchain 函数开始。在之前的实现中，NewBlockchain 会创建一个新的 Blockchain 实例，并向其中加入创世块。而现在，我们希望它做的事情有：
打开一个数据库文件
检查文件里面是否已经存储了一个区块链
如果已经存储了一个区块链：
创建一个新的 Blockchain 实例
设置 Blockchain 实例的 tip 为数据库中存储的最后一个块的哈希
如果没有区块链：
创建创世块
存储到数据库
将创世块哈希保存为最后一个块的哈希
创建一个新的 Blockchain 实例，初始时 tip 指向创世块（tip 有尾部，尖端的意思，在这里 tip 存储的是最后一个块的哈希）
*/

const dbFile = "/Users/finley/go-project/demochain/data/blockchain/blockchain.db"
const blocksBucket = "blocks"
const genesisCoinbaseData = "这是创世区块的内容"

type Blockchain struct {
	Tip []byte
	Db  *bolt.DB
}

//生成新的区块链
func NewBlockchain() *Blockchain {
	/*	gensisBlock := GenerateGenesisBlock()
		blockchain := Blockchain{}
		blockchain.ApendBlock(&gensisBlock)
		return &blockchain
	*/
	var tip []byte
	db, err := bolt.Open(dbFile, 0600, nil)

	if err != nil {
		log.Fatal(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))

		if b == nil {
			genesis := GenerateGenesisBlock()
			b, err := tx.CreateBucket([]byte(blocksBucket))
			if err != nil {
				log.Fatal(err)
			}
			err = b.Put(genesis.Hash, genesis.Serialize())
			err = b.Put([]byte("l"), genesis.Hash)
			tip = genesis.Hash
		} else {
			tip = b.Get([]byte("l"))
		}

		return nil
	})

	bc := Blockchain{tip, db}

	return &bc
}

func (bc *Blockchain) AddBlock(data string) {
	var lastHash []byte

	err := bc.Db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		lastHash = b.Get([]byte("l"))

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	newBlock := NewBlock(data, lastHash)

	err = bc.Db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))

		err := b.Put(newBlock.Hash, newBlock.Serialize())
		if err != nil {
			log.Fatal(err)
		}

		err = b.Put([]byte("l"), newBlock.Hash)
		if err != nil {
			log.Fatal(err)
		}

		bc.Tip = newBlock.Hash

		return nil
	})
}

func (bc *Blockchain) Iterator() *BlockchainIterator {
	bci := &BlockchainIterator{bc.Tip, bc.Db}

	return bci
}
