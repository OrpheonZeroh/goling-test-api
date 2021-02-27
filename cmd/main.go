package main

import (
	"golang-api/server"
	"log"
	"net/http"

)

func main() {
	server.New()
	log.Fatal(http.ListenAndServe(":8080", nil))
}