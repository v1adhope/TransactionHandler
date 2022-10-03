package verifyDataBase

import (
	"log"

	"github.com/fjl/go-couchdb"
	"github.com/v1adhope/TransactionHandler/Server/internal/configParcer"
)

func init() {
	log.Print("Please wait, initializing...")

	db, err := couchdb.NewClient(configParcer.C.Url, nil)
	if err != nil {
		log.Fatalf("couch NewClient: %v", err)
	}

	bdList, err := db.AllDBs()
	for i := range bdList {
		if bdList[i] == configParcer.C.DataBaseName {
			log.Println("Base ready to use")
			return
		}
	}
	if err != nil {
		log.Fatalf("failed to create database: %v", err)
	}

	tmp, err := db.CreateDB(configParcer.C.DataBaseName)
	if err != nil {
		log.Fatalf("failed to create database: %v", err)
	}
	log.Printf("Database successfully created \"%v\" and ready to use", tmp.Name())
}
