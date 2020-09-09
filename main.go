package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

var ver string

/*
"myapplication": [
	{
		"version": "1.0",
		"lastcommitsha": "abc57858585",
		"description" : "pre-interview technical test"
	}
]
*/

type metadata struct {
	Version       string `json:"version"`
	LastCommitSHA string `json:"lastcommitsha"`
	Description   string `json:"description"`
}

func main() {
	router := http.NewServeMux()
	router.Handle("/version", version())

	server := &http.Server{
		Addr:         ":8000",
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}

func version() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "version: %s", ver)
	})
}
