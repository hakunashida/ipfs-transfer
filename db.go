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

/*type Person struct {
	Name string
	Phone string
}*/

const (
    dialStr        = "localhost:27017"
    dbName         = "tabs"
    collectionName = "references"
    // elementsCount  = 1000
)

// var collection *mgo.Collection
var session *mgo.Session

func connectDb() {

	// HEY!
	// follow this blog post to keep sessions alive: https://medium.com/@matryer/production-ready-mongodb-in-go-for-beginners-ef6717a77219

	fmt.Println("database connection opened")

	session, err := mgo.Dial("localhost:27017")
    if err != nil {
        panic(err)
    }
    defer session.Close()

    // Optional. Switch the session to a monotonic behavior.
    session.SetMode(mgo.Monotonic, true)

    // collection := session.DB("tabs").C("references")
    // fmt.Println(collection)
    /*collection := session.DB("test").C("people")
    err = collection.Insert(&Person{"Ale", "+55 53 8116 9639"},
		&Person{"Cla", "+55 53 8402 8510"})
    if err != nil {
        log.Fatal(err)
    }

    result := Person{}
    err = collection.Find(bson.M{"name": "Ale"}).One(&result)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Phone:", result.Phone)*/
}

func collection() *mgo.Collection {
    return GetMgoSessionPerRequest().DB(dbName).C(collectionName)
}

func GetMgoSessionPerRequest() *mgo.Session {
    var sessionPerRequest *mgo.Session
    sessionPerRequest = session.Copy()
    return sessionPerRequest
}

func clearDb() {
	// collection := session.DB("tabs").C("references")
	collection().RemoveAll(nil)
}

func addReference(name string, artist string, url string, pageViews int, rating float64) {
	// collection := session.DB("tabs").C("references")
	// fmt.Println(collection())
	// fmt.Printf("name: %s, artist: %s, url: %s", name, artist, url)
	err := collection().Insert(&TabReference{name, artist, url, pageViews, rating})
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Successfully added %s", url)
	}
}