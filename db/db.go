package db

import (
	"fmt"
	"log"

	bolt "go.etcd.io/bbolt"
)

// DB defines the interface for simple, minimal interaction with the KV store
// It only allows Get and Put
type DB interface {
	Get(string) string
	Put(string, string) error
	Close()
}

// Bolt db struct
type bdb struct {
	name       string
	bucketName []byte
	db         *bolt.DB
}

// Open takes a db name returns a struct that implements DB on which Get and Put methods
// can be called. It will create and set up a bolt db in the background (or open an
// existing one) and will use bolt db functions for the getting and putting
func Open(name string) (DB, error) {
	dbStruct := bdb{
		name:       name,
		bucketName: []byte("myBucket"),
	}
	boltdb, err := bolt.Open("my.db", 0666, nil)
	if err != nil {
		log.Printf("db creation error: %s", err)
		return nil, err
	}
	err = boltdb.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(dbStruct.bucketName)
		if err != nil {
			return fmt.Errorf("bucket creation error: %s", err)
		}
		return nil
	})
	if err != nil {
		log.Print(err)
		return nil, err
	}
	dbStruct.db = boltdb
	return dbStruct, nil

}

// Close just closes the underlying bolt db. This needs to be called or risk data corruption
func (dbs bdb) Close() {
	dbs.db.Close()
}

// Get takes a key and returns the value from the bolt db. If key doesn't exist, returns empty string
func (dbs bdb) Get(key string) string {
	var value string
	dbs.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(dbs.bucketName)
		value = string(b.Get([]byte(key)))
		return nil
	})
	return value
}

// Put inserts a key and value to the bolt db.
func (dbs bdb) Put(key, value string) error {
	err := dbs.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(dbs.bucketName)
		err := b.Put([]byte(key), []byte(value))
		if err != nil {
			return fmt.Errorf("Put error: %s", err)
		}
		return nil
	})
	log.Printf("Inserted key: %s and value: %s", key, value)
	return err
}
