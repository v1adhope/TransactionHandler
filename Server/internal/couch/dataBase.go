package couch

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/fjl/go-couchdb"
)

const (
	url    = "http://admin:admin@127.0.0.1:5984/"
	dbName = "transactions"
)

type data struct {
	Time   string
	From   string
	To     string
	Amount float64
}

func NewData(time, from, to string, amount float64) data {
	return data{
		Time:   time,
		From:   from,
		To:     to,
		Amount: amount,
	}
}

func connectDB() *couchdb.Client {
	db, err := couchdb.NewClient(url, nil)
	if err != nil {
		log.Fatalf("couch NewClient: %v", err)
	}
	return db
}

func WriteDB(dt *data) {
	db := connectDB()
	uuid := getUuid()

	res, err := db.DB(dbName).Put(uuid, dt, "")
	if err != nil {
		log.Fatalf("database POST: %v", err)
	}
	log.Printf("Transaction saccessed, rev: %v", res)
}

func getUuid() string {
	resp, err := http.Get("http://admin:admin@127.0.0.1:5984/_uuids")
	if err != nil {
		log.Fatalf("failed receipt uuid: %v", err)
	}
	defer resp.Body.Close()

	result := make(map[string][]string)
	json.NewDecoder(resp.Body).Decode(&result)

	return result["uuids"][0]
}

func GetDB(count int32) string {
	db := connectDB()
	result := new(interface{})
	opts := make(couchdb.Options)
	opts["limit"] = count
	opts["include_docs"] = true
	opts["descending"] = true

	db.DB(dbName).AllDocs(&result, opts)

	//TODO body response path to docs value
	strResult := "\n"
	for i := 0; i < len((*result).(map[string]interface{})["rows"].([]interface{})); i++ {
		strResult += fmt.Sprintf("#%v Transaction %v", i+1, (*result).(map[string]interface{})["rows"].([]interface{})[i].(map[string]interface{})["doc"]) + "\n"
	}
	if strResult == "\n" {
		fmt.Println("No transactions")
		return "No transactions"
	} else {
		log.Printf("Transactions given : %v", count)
		return strResult
	}
}
