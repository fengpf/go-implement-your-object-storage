package objects

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func get(w http.ResponseWriter, r *http.Request) {
	f, e := os.Open(filePath + strings.Split(r.URL.EscapedPath(), "/")[2])
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	defer f.Close()
	io.Copy(w, f)
}
