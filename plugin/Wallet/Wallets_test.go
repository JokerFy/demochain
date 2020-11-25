package Wallet

import (
	"fmt"
	"testing"
)

const TestWalletNode = "1"

func TestGetWalletByNodeId(t *testing.T) {
	ws := Wallets{}
	wallet, err := ws.GetWalletByNodeId(TestWalletNode, "1M59oANzychihR3fXJRxQGUqXxDume2KZ2")
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(wallet)
}

func TestGetWallets(t *testing.T) {
	wallets, err := GetWallets(TestWalletNode)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(wallets)
}

func TestCreateWallet(t *testing.T) {
	ws := Wallets{}
	address := ws.CreateWallet(TestWalletNode)
	fmt.Print(address)
}

func TestCreateWallets(t *testing.T) {
	wallets, _ := NewWallets(TestWalletNode)
	fmt.Print(wallets)
}
