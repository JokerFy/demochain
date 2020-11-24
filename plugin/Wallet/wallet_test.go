package Wallet

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	randkey := GetRandomString(36)
	pubKey, priKey, _ := GenerateKey(randkey)
	fmt.Print(pubKey, priKey)
}
