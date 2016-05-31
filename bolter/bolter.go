package bolter

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

type Bolter struct {
	database *bolt.DB
}

func New() *Bolter {

	db := instance()
	b := &Bolter{database: db}
	return b

}

func instance() (database *bolt.DB) {

	db, err := bolt.Open("config.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func (b *Bolter) Create(property string) {

	b.database.Update(func(tx *bolt.Tx) error {
		tx.CreateBucketIfNotExists([]byte(property))
		return nil
	})
}

func (b *Bolter) Delete(property string) {

	b.database.Update(func(tx *bolt.Tx) error {
		tx.DeleteBucket([]byte(property))
		return nil
	})
}

func (b *Bolter) Put(property string, key string, value string) {

	b.database.Update(func(tx *bolt.Tx) error {

		propr := []byte(property)

		tx.CreateBucketIfNotExists(propr)

		b := tx.Bucket(propr)
		err := b.Put([]byte(key), []byte(value))
		return err
	})
}

func (b *Bolter) Get(property string, key string) string {

	var value string

	b.database.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(property))
		if b == nil {
			fmt.Println("Works")
			value = ""
		}
		v := b.Get([]byte(key))
		value = string(v[:])
		return nil
	})

	return value
}

func (b *Bolter) GetAll(property string) string {

	var value string

	b.database.View(func(tx *bolt.Tx) error {
		tx.CreateBucketIfNotExists([]byte(property))

		bucket := tx.Bucket([]byte(property))
		cursor := bucket.Cursor()

		for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
			value = value + string(k) + "=" + string(v) + "\n"
		}

		return nil
	})

	return value
}

func (b *Bolter) Close() {
	b.database.Close()
}
