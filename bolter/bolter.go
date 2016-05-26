package bolter

import (
	"github.com/boltdb/bolt"
	"log"
)

func Create(property string){
	db := open()
	defer db.Close()

	db.Update(func(tx *bolt.Tx) error {
		tx.CreateBucketIfNotExists([]byte(property))
		return nil
	})
}

func Put(property string, key string, value string) {

	db := open()
	defer db.Close()

	db.Update(func(tx *bolt.Tx) error {

		propr := []byte(property);

		tx.CreateBucketIfNotExists(propr)

		b := tx.Bucket(propr)
		err := b.Put([]byte(key), []byte(value))
		return err
	})
}

func Get(property string, key string) (string){

	db := open()
	defer db.Close()

	var value string

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(property))
		v := b.Get([]byte(key))
		value = string(v[:])
		return nil
	})

	return value
}


func GetAll(property string) (string){

	db := open()
	defer db.Close()

	var value string

	db.View(func(tx *bolt.Tx) error {
		tx.CreateBucketIfNotExists([]byte(property))

		bucket := tx.Bucket([]byte(property))
		log.Println(bucket)
		cursor := bucket.Cursor()

		for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
			value = value + string(k) +"="+ string(v) +"\n"
		}

		return nil
	})

	return value
}



func open() (db *bolt.DB){

	db, err := bolt.Open("config.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	return db
}


