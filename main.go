package main

import (
	"log"
	"net/http"
	"strconv"

	"assignment/pkg"
)

const DefaultServicePort = 8080

func main() {
	r := pkg.NewRouter()
	log.Printf("Listening at %d port...", DefaultServicePort)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(DefaultServicePort), r))
}
