package main

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)

type TabReference struct {
	Id 			bson.ObjectId 	`json:"_id" bson:"_id,omitempty"`
    Name 		string			`json:"name" bson:"name"`
    Artist 		string			`json:"artist" bson:"artist"`
    Url 		string			`json:"url" bson:"url"`
    PageViews 	int 			`json:"page_views" bson:"page_views"`
    Rating 		float64 		`json:"rating" bson:"rating"`
}

const (
    dialStr        = "localhost:27017"
    dbName         = "tabs"
    collectionName = "references"
)

var (
    Database *mgo.Database
    Collection *mgo.Collection
)

func connectDb() {

	fmt.Println("database connection opened")

	session, err := mgo.Dial("localhost:27017")
    if err != nil {
        panic(err)
    }

    session.SetMode(mgo.Monotonic, true)
	Database = session.DB(dbName)
	Collection = Database.C(collectionName)

	// print all results
	/*var results []TabReference
	err = Collection.Find(nil).All(&results)
	if err != nil {
	    panic(err)
	}
	fmt.Println("Results All: ", results)*/
}

func disconnectDb() {
	Database.Session.Close()
}

func clearDb() {
	Collection.RemoveAll(nil)
}

func addReference(name string, artist string, url string, pageViews int, rating float64) {

	tabReference := &TabReference{
		Id: bson.NewObjectId(),
		Name: name,
		Artist: artist,
		Url: url,
		PageViews: pageViews,
		Rating: rating,
	}

	fmt.Println(tabReference)

	err := Collection.Insert(tabReference)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Successfully added %s", url)
	}
}