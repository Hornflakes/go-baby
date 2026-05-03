package main

import (
	"fmt"
	"net/http"
	"time"
)

func ooo(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	fmt.Println("server: ooo handler started")
	defer fmt.Println("server: ooo handler ended")

	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "hello\n")
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}

func main() {
	http.HandleFunc("/ooo", ooo)
	http.ListenAndServe(":8090", nil)
}
