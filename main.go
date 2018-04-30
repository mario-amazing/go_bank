package main

import (
	"cources/homework/my_db"
	"cources/homework/server"
	"log"
	"net/http"
	// "server/logger.go"
)

func main() {
	my_db.SetupDb()

	router := server.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
