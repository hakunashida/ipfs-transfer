package main

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
    // "gopkg.in/mgo.v2/bson"
)

type TabReference struct {
    Name string
    Artist string
    Url string
    PageViews int
    Rating float64
}

const (
    dialStr        = "localhost:27017"
    dbName         = "tabs"
    collectionName = "references"
)

var (
    Database *mgo.Database
)

func connectDb() {

	fmt.Println("database connection opened")

	session, err := mgo.Dial("localhost:27017")
    if err != nil {
        panic(err)
    }

    session.SetMode(mgo.Monotonic, true)
	Database = session.DB(dbName)
}

func disconnectDb() {
	Database.Session.Close()
}

func collection() *mgo.Collection {
	return Database.C(collectionName).With( Database.Session.Copy() )
}


func clearDb() {
	collection().RemoveAll(nil)
}

func addReference(name string, artist string, url string, pageViews int, rating float64) {
	err := collection().Insert(&TabReference{name, artist, url, pageViews, rating})
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Successfully added %s", url)
	}
}