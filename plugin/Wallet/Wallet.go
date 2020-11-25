package Wallet

import (
	"bytes"
	"crypto/ecdsa"
	"demochain/core"
)

//用于生成地址的版本
const Version = byte(0x00)

//用于生成地址的校验和位数
const AddressChecksumLen = 4

type Wallet struct {
	PrivateKey ecdsa.PrivateKey
	PublicKey  []byte
}

//1.新建钱包
func NewWallet() *Wallet {
	private, public := newKeyPair()
	wallet := Wallet{private, public}

	return &wallet
}

//2.获取钱包地址 根据公钥生成地址
func (wallet *Wallet) GetAddress() []byte {

	//1.使用RIPEMD160(SHA256(PubKey)) 哈希算法，取公钥并对其哈希两次
	ripemd160Hash := Ripemd160Hash(wallet.PublicKey)
	//2.拼接版本
	version_ripemd160Hash := append([]byte{Version}, ripemd160Hash...)
	//3.两次sha256生成校验和
	checkSumBytes := CheckSum(version_ripemd160Hash)
	//4.拼接校验和
	b2 := append(version_ripemd160Hash, checkSumBytes...)

	//5.base58编码
	return core.Base58Encode(b2)
}

//5.判断地址是否有效
func (wallet *Wallet) IsValidForAddress(address []byte) bool {

	//1.base58解码地址得到版本，公钥哈希和校验位拼接的字节数组
	version_publicKey_checksumBytes := core.Base58Decode(address)
	//2.获取校验位和version_publicKeHash
	checkSumBytes := version_publicKey_checksumBytes[len(version_publicKey_checksumBytes)-AddressChecksumLen:]
	version_ripemd160 := version_publicKey_checksumBytes[:len(version_publicKey_checksumBytes)-AddressChecksumLen]

	//3.重新用解码后的version_ripemd160获得校验和
	checkSumBytesNew := CheckSum(version_ripemd160)

	//4.比较解码生成的校验和CheckSum重新计算的校验和
	if bytes.Compare(checkSumBytes, checkSumBytesNew) == 0 {

		return true
	}

	return false
}
