package main

import "./core"

func main() {
	bc := core.NewBlockchain()
	bc.SendData("Send 1 BTC to jacky")
	bc.SendData("Send 2 ETC to jacky")
	bc.SendData("Send 3 ETH to jacky")
	bc.SendData("Send 4 EOS to jacky")
	bc.Print()
}
