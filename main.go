package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	fmt.Println("Karibu :)")

	connectDb()
	beginFetching()

	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8000", router))

	defer disconnectDb()
}
