package server

import (
	"log"
	"net/http"
)

func StartServer(port string, webDir string) {
	http.Handle("/", http.FileServer(http.Dir(webDir)))

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
