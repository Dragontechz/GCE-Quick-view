package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("templates")))

	log.Println("server running ")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
