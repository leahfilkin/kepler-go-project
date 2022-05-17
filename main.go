package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"sync"
)

type Metadata struct {
	Version       string `json:"version"`
	Description   string `json:"description"`
	LastCommitSha string `json:"last_commit_sha"`
}

type countHandler struct {
	mu    sync.Mutex
	count int
}

func hello(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", `text/plain`)
	isOnlyLetters := regexp.MustCompile(`^[A-Za-z\s]+$`).MatchString
	if message := req.URL.Query().Get("msg"); message != "" {
		if !isOnlyLetters(message) {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "You cannot use integers in your message string\n")
		} else {
			// if i move this to before bad request it doesnt work :(
			w.WriteHeader(http.StatusNoContent)
			fmt.Fprintf(w, message)
		}
	} else {
		w.WriteHeader(http.StatusNoContent)
		fmt.Fprintf(w, "hello world\n")
	}
}
func health(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusNoContent)
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
	w.WriteHeader(http.StatusNoContent)
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

func (handler *countHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
	handler.mu.Lock()
	handler.count++
	handler.mu.Unlock()
	fmt.Fprintf(w, "Visited count: %d", handler.count)
}

func main() {

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/health", health)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/metadata", metadata)
	http.Handle("/count", new(countHandler))

	http.ListenAndServe(":8090", nil)
}

// another handler called counter
// handler with state
// http package - handle, struct that implements interface, with counter
// only one request increment at a time
// why struct over global state - reading
// testing that only writing one thing at a time
// mutexes
