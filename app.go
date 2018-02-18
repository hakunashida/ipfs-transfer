package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

func main() {

	fmt.Println("Ushirikina")

	args := os.Args[1:]

	connectDb()

	if contains(args, "crawl") || contains(args, "C") {
		beginFetching()
	}

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	router := NewRouter()

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable was not set")
	}

	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(originsOk, headersOk, methodsOk)(router)))

	defer disconnectDb()
}