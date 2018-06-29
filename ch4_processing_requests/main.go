package main

import (
	"net/http"
	"fmt"
)

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/headers", headers)
	server.ListenAndServe()
}
func headers(w http.ResponseWriter, r *http.Request) {
	h := r.Header
	fmt.Fprintln(w, h)
}
