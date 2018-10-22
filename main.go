package main

import "github.com/bytepower/blockchain/blockchain"

func main() {
	nbc := blockchain.NewBlockchain()
	defer nbc.DB.Close()

	cli := blockchain.CLI{BC: nbc}
	cli.Run()
}
