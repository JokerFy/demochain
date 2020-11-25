package core

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"math/big"
)

/* hash函数
params: available_str string ---> pre hash blockchain
*/
func HashString(available_str string) string {
	h := sha256.New()
	h.Write([]byte(available_str))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buff.Bytes()
}

func DataToHash(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[:]
}

//base64:ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/
//base58:去掉0(零)，O(大写的 o)，I(大写的i)，l(小写的 L)，+，/

//base58编码集
var b58Alphabet = []byte("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")

// 字节数组转 Base58,加密
func Base58Encode(input []byte) []byte {

	var result []byte

	x := big.NewInt(0).SetBytes(input)

	base := big.NewInt(int64(len(b58Alphabet)))
	zero := big.NewInt(0)
	mod := &big.Int{}

	for x.Cmp(zero) != 0 {

		x.DivMod(x, base, mod)
		result = append(result, b58Alphabet[mod.Int64()])
	}

	ReverseBytes(result)
	for b := range input {

		if b == 0x00 {

			result = append([]byte{b58Alphabet[0]}, result...)
		} else {

			break
		}
	}

	return result
}

// Base58转字节数组，解密
func Base58Decode(input []byte) []byte {

	result := big.NewInt(0)
	zeroBytes := 0

	for b := range input {

		if b == 0x00 {

			zeroBytes++
		}
	}

	payload := input[zeroBytes:]
	for _, b := range payload {

		charIndex := bytes.IndexByte(b58Alphabet, b)
		result.Mul(result, big.NewInt(58))
		result.Add(result, big.NewInt(int64(charIndex)))
	}

	decoded := result.Bytes()
	//decoded...表示将decoded所有字节追加
	decoded = append(bytes.Repeat([]byte{byte(0x00)}, zeroBytes), decoded...)

	return decoded
}

// ReverseBytes reverses a byte array
func ReverseBytes(data []byte) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}
