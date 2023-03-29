package main

import (
	"log"
	"net/http"

	"github.com/qbradq/petshop-go/data"
)

// staticd is the static asset and template service for Pet Shop.
func main() {
	hfs := http.FileServer(http.FS(data.StaticFS))
	http.Handle("/", hfs)
	if err := http.ListenAndServe("0.0.0.0:8080", nil); err != nil {
		log.Println(err)
	}
}
