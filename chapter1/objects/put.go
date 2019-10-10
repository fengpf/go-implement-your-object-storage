package objects

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

var (
	filePath = "./storage/"
)

func put(w http.ResponseWriter, r *http.Request) {
	f, e := os.Create(filePath + strings.Split(r.URL.EscapedPath(), "/")[2])
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer f.Close()
	io.Copy(f, r.Body)
}
