package core

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Block struct {
	Index         int64  //区块编号
	Timestamp     int64  //区块时间戳
	PrevBlockHash string //上一个区块哈希值
	Hash          string //当前区块哈希值
	Data          string //区块数据

}

func calculateHash(b Block) string {
	blockDdata := string(b.Index) + string(b.Timestamp) + string(b.PrevBlockHash) + string(b.Data)
	hashInBytes := sha256.Sum256([]byte(blockDdata))
	return hex.EncodeToString(hashInBytes[:])
}

func GenereateNewBlock(preBlock Block, data string) Block {
	newBlock := Block{}
	newBlock.Index = preBlock.Index + 1
	newBlock.PrevBlockHash = preBlock.Hash
	newBlock.Timestamp = time.Now().Unix()
	newBlock.Data = data
	newBlock.Hash = calculateHash(newBlock)
	return newBlock
}

func GenerateGenesisBlock() Block {
	preBlock := Block{}
	preBlock.Index = -1
	preBlock.Hash = ""
	return GenereateNewBlock(preBlock, "Genesis Block")
}
