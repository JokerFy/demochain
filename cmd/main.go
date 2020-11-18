package main

import "demochain/core"

func main() {
	bc := core.NewBlockchain()
	bc.SendDdata("Send 1 BTC to Jacky")
	bc.SendDdata("Send 1 EOS to Jacky")
	bc.Print()
}
