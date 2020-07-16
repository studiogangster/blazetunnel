package db

import (
	"errors"
	"fmt"
	"log"
	"sync"

	"github.com/boltdb/bolt"
)

var appsBucket = "apps"
var db *bolt.DB

var doOnce sync.Once

func initDb() {

	_db, err := bolt.Open("tmp.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	// defer db.Close()

	db = _db

	// Create users bucket
	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(appsBucket))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})
}

func initialize() {

	doOnce.Do(func() {
		initDb()
	})
}

func createapp(appname string, password string) error {

	initialize()
	if appname == "" {
		return errors.New("Invalid appname")
	}

	if password == "" {
		return errors.New("Invalid password")
	}

	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(appsBucket))
		exists := string(b.Get([]byte(appname)))
		if exists != "" {
			return errors.New("Appname already exists")
		}

		err := b.Put([]byte(appname), []byte(password))

		return err
	})

	return err

}

func getapp(appname string) (string, error) {
	initialize()
	var value = ""

	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(appsBucket))

		value = string(b.Get([]byte(appname)))

		return nil
	})
	return value, err
}

func main() {

	log.Println("Creating AppI")
	createapp("test_app", "test_password")
	log.Println("Reading AppI")

	getapp("test_app")
	getapp("232")

}
