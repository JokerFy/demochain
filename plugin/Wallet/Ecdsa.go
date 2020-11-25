package Wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"golang.org/x/crypto/ripemd160"
	"log"
)

//通过私钥创建公钥
func newKeyPair() (ecdsa.PrivateKey, []byte) {

	//1.椭圆曲线算法生成私钥
	curve := elliptic.P256()
	privateKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {

		log.Panic(err)
	}

	//2.通过私钥生成公钥
	publicKey := append(privateKey.PublicKey.X.Bytes(), privateKey.PublicKey.Y.Bytes()...)

	return *privateKey, publicKey
}

//将公钥进行两次哈希
func Ripemd160Hash(publicKey []byte) []byte {

	//1.hash256
	hash256 := sha256.New()
	hash256.Write(publicKey)
	hash := hash256.Sum(nil)

	//2.ripemd160
	ripemd160 := ripemd160.New()
	ripemd160.Write(hash)

	return ripemd160.Sum(nil)
}

//两次sha256哈希生成校验和
func CheckSum(bytes []byte) []byte {

	//hasher := sha256.New()
	//hasher.Write(bytes)
	//hash := hasher.Sum(nil)
	//与下面一句等同
	//hash := sha256.Sum256(bytes)

	hash1 := sha256.Sum256(bytes)
	hash2 := sha256.Sum256(hash1[:])

	return hash2[:AddressChecksumLen]
}
