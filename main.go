package main

import (
	"github.com/nomadcoders/nomadcoin/cli"
	"github.com/nomadcoders/nomadcoin/db"
)

func main() {
	// defer 은 함수가 종료시 실행
	defer db.Close()
	db.InitDB()

	// Blockchain 을 두번 쓰면 데드락 걸림
	// blockchain.Blockchain()
	cli.Start()
}
