package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var (
	httpPort    int
	commit      string
	description string
	httpAddr    string
	logFormat   string
	version     string
)

type metadata struct {
	Version       string `json:"version"`
	LastCommitSHA string `json:"lastcommitsha"`
	Description   string `json:"description"`
}

type requestLog struct {
	RemoteHost string `json:"host"`
	Identity   string `json:"identity"`
	User       string `json:"user"`
	Date       string `json:"timestamp"`
	Method     string `json:"method"`
	Path       string `json:"path"`
	Protocol   string `json:"protocol"`
	Status     int    `json:"status"`
	Size       int    `json:"contentlength"`
	Referer    string `json:"referer"`
	UserAgent  string `json:"useragent"`
}

func main() {
	flag.StringVar(&logFormat, "log", "kv", "the log format to use [none|combined|json|kv]")
	flag.StringVar(&httpAddr, "address", "127.0.0.1", "The TCP address to listen on")
	flag.IntVar(&httpPort, "port", 8000, "The TCP port to listen on")
	flag.Parse()

	d, exists := os.LookupEnv("DESCRIPTION")
	if !exists {
		d = "pre-interview technical test"
	}
	description = d

	addr := fmt.Sprintf("%s:%d", httpAddr, httpPort)

	handler := http.NewServeMux()
	handler.HandleFunc("/version", versionHandler)

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

func versionHandler(w http.ResponseWriter, r *http.Request) {
	metadata, err := formatVersion(version, commit, description)
	if err != nil {
		fmt.Fprintf(w, "Service Error")
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}
	response := fmt.Sprintf("\"myapplication\": %s\n", metadata)
	logRequest(r, http.StatusOK, len(response))
	fmt.Fprint(w, response)
}

func formatVersion(v string, c string, d string) (string, error) {
	m := metadata{
		Version:       v,
		LastCommitSHA: c,
		Description:   d,
	}

	ma := []metadata{m}

	mb, err := json.MarshalIndent(ma, "", "  ")
	if err != nil {
		log.Println(err.Error())
		return "", err
	}

	return string(mb), nil
}

func logRequest(r *http.Request, status int, size int) {

	if logFormat == "none" {
		return
	}

	var out string

	t := time.Now()
	u, _, ok := r.BasicAuth()
	if !ok {
		u = "-"
	}
	a := strings.Split(r.RemoteAddr, ":")

	lr := requestLog{
		RemoteHost: a[0],
		Identity:   "-",
		User:       u,
		Date:       t.Format(time.RFC3339),
		Method:     r.Method,
		Path:       r.URL.Path,
		Protocol:   r.Proto,
		Status:     status,
		Size:       size,
		UserAgent:  r.UserAgent(),
		Referer:    r.Referer(),
	}

	switch logFormat {
	case "combined":
		out = fmt.Sprintf("%s - %s [%s] \"%s %s %s\" %d %d \"%s\" \"%s\"",
			a[0], u, t.Format("2/01/2006:15:04:05 -0700"), r.Method, r.URL.Path,
			r.Proto, status, size, r.Referer(), r.UserAgent(),
		)
	case "kv":
		out = fmt.Sprintf("%+v", lr)
		out = out[1 : len(out)-1] // remove {} from printed struct
	case "json":
		b, err := json.Marshal(lr)
		if err != nil {
			log.Print(err.Error())
		}
		out = string(b)
	default:
		out = fmt.Sprintf("%+v", lr)
		out = out[1 : len(out)-1]
	}

	fmt.Println(out)
}
