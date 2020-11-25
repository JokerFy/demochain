package Wallet

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	w := NewWallet()
	fmt.Printf("Your new addres：%s\n", w.GetAddress())
}

func TestWalletValidate(t *testing.T) {
	w := Wallet{}
	address := "1PnN9R2PWyvNaPTZABs5MbdmGnWwUtE6Vh"
	fmt.Print(w.IsValidForAddress([]byte(address)))
}
