package main

import (
	"log"
	"net/http"

	"go-object-storage/chapter2/dataServer/heartbeat"
	"go-object-storage/chapter2/dataServer/locate"
	"go-object-storage/chapter2/dataServer/objects"
)

func main() {
	go heartbeat.StartHeartbeat()
	go locate.StartLocate()
	http.HandleFunc("/objects/", objects.Handler)
	log.Fatal(http.ListenAndServe("9000", nil))
}
