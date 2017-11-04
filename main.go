package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	/*hash := ipfsSave("test3")
	fmt.Println(hash)
	fmt.Println(ipfsLoad("QmbpoQhY8G2NXER8JFxcLKei6QTPZtRc3sayayoFw8TsEk"))*/

	fmt.Println("Karibu :)")

	connectDb()
	beginFetching()

	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8000", router))

	defer disconnectDb()
}
