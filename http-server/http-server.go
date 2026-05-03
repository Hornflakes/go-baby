package main

import (
	"fmt"
	"net/http"
)

func eee(w http.ResponseWriter, _req *http.Request) {
	fmt.Fprintf(w, "eee\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	http.HandleFunc("/eee", eee)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":8090", nil)
}
