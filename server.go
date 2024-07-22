package main

import (
	"libary/router"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe("localhost:4000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	router.Routes(w, r)
}
