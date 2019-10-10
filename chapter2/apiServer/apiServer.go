package main

import (
	"log"
	"net/http"

	"go-object-storage/chapter2/apiServer/heartbeat"
	"go-object-storage/chapter2/apiServer/locate"
	"go-object-storage/chapter2/apiServer/objects"
)

func main() {
	go heartbeat.ListenHeartbeat()
	http.HandleFunc("/objects/", objects.Handler)
	http.HandleFunc("/locate/", locate.Handler)
	log.Fatal(http.ListenAndServe("9000", nil))
}
