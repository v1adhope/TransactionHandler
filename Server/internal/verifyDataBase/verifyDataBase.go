package verifyDataBase

import (
	"log"

	"github.com/fjl/go-couchdb"
)

const (
	url    = "http://admin:admin@127.0.0.1:5984/"
	dbName = "transactions"
)

func init() {
	log.Print("Please wait, initializing...")

	db, err := couchdb.NewClient(url, nil)
	if err != nil {
		log.Fatalf("couch NewClient: %v", err)
	}

	bdList, err := db.AllDBs()
	for i := range bdList {
		if bdList[i] == "transactions" {
			log.Println("Base ready to use")
			return
		}
	}
	if err != nil {
		log.Fatalf("failed to create database: %v", err)
	}

	tmp, err := db.CreateDB(dbName)
	if err != nil {
		log.Fatalf("failed to create database: %v", err)
	}
	log.Printf("Database successfully created \"%v\" and ready to use", tmp.Name())
}
