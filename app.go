package main

import (
	"fmt"
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
	w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")

	host := r.Host

	if host == "" || host == "/" {
		log.Print("Host is not valid: " + host)
		errorHandler(w, r, http.StatusNotFound)
		return
	}

	prefix := "/files/"
	urlPath := r.URL.Path
	path := prefix + host + urlPath

	if urlPath == "/index.html" {
		log.Print("Request index: " + path)

		f, err := os.Open(path)
		if err != nil {
			errorHandler(w, r, http.StatusNotFound)
			return
		}
		http.ServeContent(w, r, path, time.Now(), f)
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		errorHandler(w, r, http.StatusNotFound)
		return
	}

	log.Print("Request: " + path)
	http.ServeFile(w, r, path)
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "")
	}
}
