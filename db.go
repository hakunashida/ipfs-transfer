package main

import (
	"fmt"
	"log"
	// "context"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
    elastic "gopkg.in/olivere/elastic.v5"
)

type TabReference struct {
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
    Client *elastic.Client
)

func connectDb() {

	fmt.Println("database connection opened")

	/*
	*	Set up mongo
	**/
	session, err := mgo.Dial("localhost:27017")
    if err != nil {
        panic(err)
    }

    session.SetMode(mgo.Monotonic, true)
	Database = session.DB(dbName)
	Collection = Database.C(collectionName)

	// clearDb()
	printDb()

	/*
	*	Set up elastic
	**/

	// Create a client
	Client, err = elastic.NewClient()
	if err != nil {
		panic(err)
	}

	// Create an index
	/*_, err = Client.CreateIndex("tabs").Do(context.Background())
	if err != nil {
		// Handle error
		panic(err)
	}*/
}

func disconnectDb() {
	Database.Session.Close()
}

func printDb() {
	// print all results
	var results []TabReference
	err := Collection.Find(nil).All(&results)
	if err != nil {
	    panic(err)
	}
	fmt.Println("All results: ", results)
}

func clearDb() {
	Collection.RemoveAll(nil)
}

func addReference(name string, artist string, url string, pageViews int, rating float64) {

	// Use the url as a unique identifier to determine if a new record should be inserted
	resultsCount, err := Collection.Find(bson.M{"url": url}).Count()
	if err != nil {
	    panic(err)
	}
	
	// Don't insert if the record already exists
	if (resultsCount > 0) {
		fmt.Println("Skipping " + url + " because it has already been added")
	} 

	// Otherwise, insert a new record
	else {

		tabReference := &TabReference{
			Name: name,
			Artist: artist,
			Url: url,
			PageViews: pageViews,
			Rating: rating,
		}

		err := Collection.Insert(tabReference)

		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Printf("Successfully added %s", url)
		}
	}


	// Add a document to the index
	/*_, err = Client.Index().
		Index("tabs").
		Type("tabReference").
		BodyJson(tabReference).
		Refresh("true").
		Do(context.Background())
	if err != nil {
		// Handle error
		panic(err)
	}

	// Search with a term query
	termQuery := elastic.NewTermQuery("name", "my")
	searchResult, err := Client.Search().
	    Index("tabs").   // search in index "twitter"
	    Query(termQuery).   // specify the query
	    Sort("name", true). // sort by "user" field, ascending
	    // From(0).Size(10).   // take documents 0-9
	    Pretty(true).       // pretty print request and response JSON
	    Do(context.Background())             // execute
	if err != nil {
	    // Handle error
	    panic(err)
	}

	fmt.Println(searchResult)*/
}