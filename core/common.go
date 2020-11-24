package core

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
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
