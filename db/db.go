package db

import (
	"fmt"
	"github.com/nomadcoders/nomadcoin/utils"
	bolt "go.etcd.io/bbolt"
	"os"
)

const (
	dbName      = "blockchain"
	dataBucket  = "data"
	blockBucket = "blocks"

	checkpoint = "checkpoint"
)

var db *bolt.DB

type DB struct {
}

// blockchain 용이 아닌 테스트 용
func getDbName() string {
	port := os.Args[2][6:]
	return fmt.Sprintf("%s_%s.db", dbName, port)
}

func (DB) FindBlock(hash string) []byte {
	return findBlock(hash)
}
func (DB) LoadChain() []byte {
	return loadChain()
}
func (DB) SaveBlock(hash string, data []byte) {
	saveBlock(hash, data)
}
func (DB) SaveChain(data []byte) {
	saveChain(data)
}
func (DB) DeleteAllBlocks() {
	deleteAllBlocks()
}

func InitDB() {
	if db == nil {
		dbPointer, err := bolt.Open(getDbName(), 0600, nil)
		db = dbPointer
		utils.HandleErr(err)
		err = db.Update(func(tx *bolt.Tx) error {
			_, err := tx.CreateBucketIfNotExists([]byte(dataBucket))
			utils.HandleErr(err)
			_, err = tx.CreateBucketIfNotExists([]byte(blockBucket))
			return err
		})
		utils.HandleErr(err)

		// bolt.bucket = mysql.table
	}
}

func Close() {
	db.Close()
}

func saveBlock(hash string, data []byte) {
	err := db.Update(func(t *bolt.Tx) error {
		bucket := t.Bucket([]byte(blockBucket))
		err := bucket.Put([]byte(hash), data)
		return err
	})
	utils.HandleErr(err)
}

func saveChain(data []byte) {
	err := db.Update(func(t *bolt.Tx) error {
		bucket := t.Bucket([]byte(dataBucket))
		err := bucket.Put([]byte(checkpoint), data)
		return err
	})
	utils.HandleErr(err)
}

// loadChain : DB에서 마지막 체크포인트 가져오기
func loadChain() []byte {
	var data []byte
	utils.HandleErr(db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(dataBucket))
		data = bucket.Get([]byte(checkpoint))
		return nil
	}))
	return data
}

// getBlock : DB에서 블록 값 가져오기
func findBlock(hash string) []byte {
	var data []byte
	utils.HandleErr(db.View(func(t *bolt.Tx) error {
		bucket := t.Bucket([]byte(blockBucket))
		data = bucket.Get([]byte(hash))
		return nil
	}))
	return data
}

func deleteAllBlocks() {
	utils.HandleErr(db.Update(func(t *bolt.Tx) error {
		utils.HandleErr(t.DeleteBucket([]byte(blockBucket)))
		_, err := t.CreateBucket([]byte(blockBucket))
		utils.HandleErr(err)
		return nil
	}))
}
