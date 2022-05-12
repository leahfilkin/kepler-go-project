package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	// return content type, text plain - header
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

func main() {

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/health", health)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":8090", nil)
}
