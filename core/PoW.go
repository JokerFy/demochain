package core

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math"
	"math/big"
	"time"
)

type ProofOfWork struct {
	block  *Block   // 即将生成的区块对象
	target *big.Int //生成区块的难度
}

const targetBits = 20

func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	//难度：target=10的18次方（即要求计算出的hash值小于这个target）
	target.Lsh(target, uint(256-targetBits))
	pow := &ProofOfWork{b, target}
	return pow
}

func (pow *ProofOfWork) prepareData(nonce int64) string {
	data := string(pow.block.Index) + string(pow.block.Timestamp) + string(pow.block.PrevBlockHash) + string(pow.block.Data) + string(nonce)
	hashInBytes := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hashInBytes[:])
}

func (pow *ProofOfWork) Run() (int64, string) {
	nonce := 0
	t1 := time.Now().UnixNano()
	var hashVar string
	for nonce < math.MaxInt64 {
		data := pow.prepareData(int64(nonce))

		if data[0:4] == "0000" {
			hashVar = data
			break
		} else {
			nonce++
		}
	}

	fmt.Println(nonce)
	fmt.Printf("PoW Duration:%.3fs\n", float64((time.Now().UnixNano()-t1)/1e6)/1000)

	fmt.Printf("\n\n")

	return int64(nonce), hashVar
}
