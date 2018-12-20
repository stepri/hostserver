package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	log.Print("Server started!")
	http.ListenAndServe(":7777", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	host := r.Host

	if host == "" || host == "/" {
		log.Print("Host is not valid: " + host)
		http.NotFound(w, r)
		return
	}

	prefix := "/files/"
	urlPath := r.URL.Path
	path := prefix + host + urlPath

	if urlPath == "/index.html" {
		log.Print("Request index: " + path)

		f, _ := os.Open(path)
		http.ServeContent(w, r, path, time.Now(), f)
	}

	log.Print("Request: " + path)
	http.ServeFile(w, r, path)
}
