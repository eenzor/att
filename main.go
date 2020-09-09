package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

var ver string
var commit string

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
	handler := http.NewServeMux()
	handler.Handle("/version", version())

	server := &http.Server{
		Addr:         ":8000",
		Handler:      handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}

func version() http.Handler {

	m := metadata{
		Version:       ver,
		LastCommitSHA: commit,
		Description:   "pre-interview technical test",
	}

	ma := []metadata{m}

	mb, err := json.MarshalIndent(ma, "", "  ")
	if err != nil {
		fmt.Println(err.Error())
	}

	response := fmt.Sprintf("\"myapplication\": %s\n", string(mb))

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, response)
	})
}
