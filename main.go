package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
)

func main() {

	fmt.Println("Ushirikina")

	connectDb()
	// beginFetching()

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(originsOk, headersOk, methodsOk)(router)))

	defer disconnectDb()
}
