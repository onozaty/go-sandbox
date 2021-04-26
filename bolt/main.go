package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

func main() {

	flag.Parse()
	args := flag.Args()

	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("MyBucket"))
		if err != nil {
			return err
		}

		if len(args) == 1 {
			v := b.Get([]byte(args[0]))
			fmt.Printf("Get: key(%s) value(%s)\n", args[0], v)
		} else if len(args) == 2 {
			err := b.Put([]byte(args[0]), []byte(args[1]))
			if err != nil {
				return err
			}
			fmt.Printf("Put: id(%s) value(%s)\n", args[0], args[1])
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
}
