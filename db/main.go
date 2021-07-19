package  main

import (
	"github.com/boltdb/bolt"
	"log"
	"time"
)

func main(){
	db, err := bolt.Open("my.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}