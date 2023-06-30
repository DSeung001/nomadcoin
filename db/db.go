package db

import (
	"github.com/boltdb/bolt"
	"github.com/nomadcoders/nomadcoin/utils"
)

const (
	dbName      = "blockchain.db"
	dataBucket  = "data"
	blockBucket = "blocks"
)

var db *bolt.DB

func DB() *bolt.DB {
	if db == nil {
		dbPointer, err := bolt.Open(dbName, 0600, nil)
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
	return db
}
