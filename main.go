package main

import (
	"github.com/nomadcoders/nomadcoin/blockchain"
	"github.com/nomadcoders/nomadcoin/cli"
	"github.com/nomadcoders/nomadcoin/db"
)

func main() {
	// defer 은 함수가 종료시 실행
	defer db.Close()

	blockchain.Blockchain()
	cli.Start()
}
