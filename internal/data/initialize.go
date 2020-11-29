package data

import (
	"log"

	"github.com/dgraph-io/badger/v2"
)

func InitializeData() {
	db, err := badger.Open(badger.DefaultOptions("./matros"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
