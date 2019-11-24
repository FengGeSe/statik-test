//go:generate statik -src=./static -f

package main

import (
	"log"
	"net/http"

	"github.com/rakyll/statik/fs"

	_ "statik-test/statik"
)

func main() {
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}

	// Serve the contents over HTTP.
	http.Handle("/", http.StripPrefix("/", http.FileServer(statikFS)))

	log.Println("file server is running 127.0.0.1:8080")
	http.ListenAndServe(":8080", nil)
}
