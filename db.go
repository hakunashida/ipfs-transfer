package main

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	elastic "gopkg.in/olivere/elastic.v5"
)

const (
	dialStr        = "localhost:27017"
	dbName         = "tabs"
	collectionName = "references"
)

var (
	Database   *mgo.Database
	Collection *mgo.Collection
	Client     *elastic.Client
)

func connectDb() {

	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}

	fmt.Println("1. Database connection opened")

	session.SetMode(mgo.Monotonic, true)
	Database = session.DB(dbName)
	Collection = Database.C(collectionName)

	// clearDb()

	// set up search indexes
	err = Collection.EnsureIndexKey("name", "artist")
	if err != nil {
		panic(err)
	}
}

func disconnectDb() {
	Database.Session.Close()
}

func getAllTabs() []TabReference {

	// print all database records
	var results []TabReference
	err := Collection.Find(nil).All(&results)
	if err != nil {
		panic(err)
	}
	return results
}

func getTabById(id string) TabReference {
	data := TabReference{}
	err := Collection.FindId(bson.ObjectIdHex(id)).One(&data)
	if err != nil {
		panic(err)
	}
	return data
}

func clearDb() {
	Collection.RemoveAll(nil)
}

func searchDb(terms string) TabReferences {

	var tabReferences TabReferences
	query := bson.M{"$text": bson.M{"$search": terms}}
	err := Collection.Find(query).All(&tabReferences)
	if err != nil {
		panic(err)
	}

	log.Printf("RunQuery : %d : Count[%d]\n", query, len(tabReferences))
	return tabReferences
}

func addReference(name string, artist string, url string, ipfsHash string, pageViews int, rating float64) {

	// Use the url as a unique identifier to determine if a new record should be inserted
	resultsCount, err := Collection.Find(bson.M{"url": url}).Count()
	if err != nil {
		panic(err)
	}

	if resultsCount > 0 {

		// Don't insert if the record already exists
		fmt.Println("Skipping " + url + " because it has already been added")
	} else {

		// Otherwise, insert a new record
		tabReference := &TabReference{
			Name:      name,
			Artist:    artist,
			Url:       url,
			PageViews: pageViews,
			Rating:    rating,
			IpfsHash:  ipfsHash,
		}

		err := Collection.Insert(tabReference)

		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Printf("Successfully added %s", url)
		}
	}
}
