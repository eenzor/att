package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

var (
	version  string
	commit   string
	logJSON  bool
	httpAddr string
	httpPort int
)

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

type requestLog struct {
	Date     string `json:"timestamp"`
	Host     string `json:"host"`
	Method   string `json:"method"`
	Path     string `json:"path"`
	Protocol string `json:"protocol"`
}

func main() {
	flag.BoolVar(&logJSON, "json", false, "log ouput as JSON")
	flag.StringVar(&httpAddr, "address", "", "The TCP address to listen on")
	flag.IntVar(&httpPort, "port", 8000, "The TCP port to listen on")
	flag.Parse()

	addr := fmt.Sprintf("%s:%d", httpAddr, httpPort)

	handler := http.NewServeMux()
	handler.Handle("/version", handleVersion())

	server := &http.Server{
		Addr:         addr,
		Handler:      handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	log.Printf("Listening on %s\n", addr)
	log.Fatal(server.ListenAndServe())
}

func handleVersion() http.Handler {

	response := fmt.Sprintf("\"myapplication\": %s\n", formatVersion(version, commit))

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logRequest(r)
		fmt.Fprintf(w, response)
	})
}

func formatVersion(v string, c string) string {
	m := metadata{
		Version:       v,
		LastCommitSHA: c,
		Description:   "pre-interview technical test",
	}

	ma := []metadata{m}

	mb, err := json.MarshalIndent(ma, "", "  ")
	if err != nil {
		fmt.Println(err.Error())
	}

	return string(mb)
}

func logRequest(r *http.Request) {
	lr := requestLog{
		Host:     r.RemoteAddr,
		Method:   r.Method,
		Path:     r.URL.Path,
		Protocol: r.Proto,
		Date:     time.Now().Format(time.RFC3339),
	}

	out := fmt.Sprintf("%+v", lr)

	if logJSON {
		b, err := json.Marshal(lr)
		if err != nil {
			log.Printf(err.Error())
		}
		out = string(b)
	}

	fmt.Println(out)
}
