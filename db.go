package main

import (
	"fmt"
	"log"
	"os"

	"github.com/boltdb/bolt"
)

func createDB() *bolt.DB {
	dataLoc := "./data"
	path := dataLoc + "/.bloke"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 755)
	}
	fmt.Println("Path: ", path)
	db, err := bolt.Open(path+"/bolt.db", 0644, nil)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func add(key, value []byte) {
	db := createDB()
	defer db.Close()

	err := db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte("bloke"))
		if err != nil {
			return err
		}
		err = bucket.Put(key, value)
		if err != nil {
			return err
		}
		fmt.Printf("%s was added to the db\n", string(key))
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
}
