package main

import (
	"net/http"
	"fmt"
	json2 "encoding/json"
)

func main() {
	server := http.Server{
		Addr:"127.0.0.1:8080",
	}
	http.HandleFunc("/write", writeExample)
	http.HandleFunc("/writeheader", writeHeaderExample)
	http.HandleFunc("/redirect", headerExample)
	http.HandleFunc("/json", headerJSONExmaple)
	server.ListenAndServe()
}
func headerJSONExmaple(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	post := &struct {
		User string
		Threads []string
	}{
		User: "Kengo Suzuki",
		Threads: []string{"first", "second", "third"},
	}
	json, _ :=  json2.Marshal(post)
	w.Write(json)
}

func headerExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "http://google.com")
	// WriteHeader prevents the header from being modified after it's called.
	w.WriteHeader(302)
}

func writeHeaderExample(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	// WriteHeader prevents the header from being modified after it's called.
	fmt.Fprintln(w, "No such service, try next door")
}


func writeExample(w http.ResponseWriter, r *http.Request) {
	str := `<html>
        <head><title>Go Web Programming</title></head>
        <body><h1>Hello World</h1></body>
        </html>`

    // w.WriteHeader allows you to set response code
    // if w.WriteHeader is not set, when w.Write is called 200 is set by default.
    w.Write([]byte(str))
}
