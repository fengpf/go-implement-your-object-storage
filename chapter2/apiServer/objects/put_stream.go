package objects

import (
	"fmt"

	"go-object-storage/chapter2/apiServer/heartbeat"
	"go-object-storage/src/lib/objectstream"
)

func putStream(object string) (*objectstream.PutStream, error) {
	server := heartbeat.ChooseRandomDataServer()
	if server == "" {
		return nil, fmt.Errorf("cannot find any dataServer")
	}

	return objectstream.NewPutStream(server, object), nil
}
