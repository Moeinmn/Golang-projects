package db

import (
	"github.com/boltdb/bolt"
	"time"
)

var taskBucket = []byte("tasks")

var db *bolt.DB

func Init(dbPath string) error {
	var err error
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		panic(err)
	}
	defer db.Close()

	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskBucket)
		return err
	})
}
