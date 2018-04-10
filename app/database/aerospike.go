package database

import (
	"github.com/aerospike/aerospike-client-go"
	"log"
)


func OpenAerospike() (*aerospike.Client){
	db, err := aerospike.NewClient("localhost", 3000)
	if err != nil {
		log.Println(err)
	}
	return db
}
