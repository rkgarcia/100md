package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	db := getDB()
	defer db.Close()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	listenAddr := ":" + port
	log.Printf("Listening on %s...\n", listenAddr)
	http.ListenAndServe(listenAddr, routes())
}
