package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	connectDb()
	// beginFetching()

	fmt.Println("Karibu :)")

	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8000", router))

	defer disconnectDb()
}
