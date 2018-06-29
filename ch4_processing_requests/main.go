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
	http.HandleFunc("/body", body)
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
func process(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprintln(w, r.Form)
}

// curl -id "first_name=sausheong&last_name=chang" 127.0.0.1:8080/body
func body(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength
	body := make([]byte, len)
	// you can call on the Read and Close methods of the Body field.
	r.Body.Read(body)
	fmt.Fprintln(w, string(body))
}

func headers(w http.ResponseWriter, r *http.Request) {
	h := r.Header
	fmt.Println(r.Header["Accept-Encoding"])
	fmt.Println(r.Header.Get("Accept-Encoding"))
	fmt.Fprintln(w, h)
}

