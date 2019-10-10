package main

import (
	"log"
	"net/http"

	"go-object-storage/chapter1/objects"
)

func main() {
	http.HandleFunc("/objects/", objects.Handler)
	log.Fatal(http.ListenAndServe(":9000", nil))
}
