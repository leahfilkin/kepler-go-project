package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Metadata struct {
	Version       string `json:"version"`
	Description   string `json:"description"`
	LastCommitSha string `json:"last_commit_sha"`
}

func hello(w http.ResponseWriter, req *http.Request) {
	// return content type, text plain - header
	w.Header().Add("Content-Type", `text/plain`)
	fmt.Fprintf(w, "hello world\n")
}

func health(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Status code: %s", http.StatusText(http.StatusNoContent))
}
func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func metadata(w http.ResponseWriter, req *http.Request) {
	myapplication := make(map[string][]Metadata)
	metadata := Metadata{
		Version:       "1.0",
		Description:   "pre-interview technical test",
		LastCommitSha: "abc57858585",
	}
	myapplication["myapplication"] = []Metadata{metadata}
	//	myapplication["myapplication"] = append(myapplication["myapplication"], metadata)
	enc := json.NewEncoder(w)
	w.Header().Add("Content-Type", `application/json`)
	enc.SetIndent("", "    ")
	if err := enc.Encode(myapplication); err != nil {
		fmt.Fprintf(w, "Error: %s", err)
	}
}
func main() {

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/health", health)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/metadata", metadata)

	http.ListenAndServe(":8090", nil)
}
