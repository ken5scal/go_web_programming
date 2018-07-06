package main

import (
	"net/http"
	"strconv"
	"path"
	"encoding/json"
)

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/post", handleRequest)
	server.ListenAndServe()
}
func handleRequest(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = handleGet(w, r)
	case "POST":
		err = handlePost(w, r)
	case "PUT":
		err = handlePut(w, r)
	case "DELETE":
		err = handleDelete(w,r)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func handleGet(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}

	post, err := retrieve(id)
	if err != nil {
		return
	}
	output, err := json.MarshalIndent(&post, "", "\t\t")
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

func handlePost(w http.ResponseWriter, r *http.Request) (err error) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)

	var post Post
	json.Unmarshal(body, &post)
	if err = post.Create(); err !=nil{
		return
	}
	w.WriteHeader(200)
	return
}

func handlePut(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}

	post, err := retrieve(id)
	if err != nil {
		return
	}

	body := make([]byte, r.ContentLength)
	r.Body.Read(body)

	json.Unmarshal(body, &post)
	if err = post.Update(); err != nil {
		return
	}
	w.WriteHeader(200)
	return
}

func handleDelete(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}

	post, err := retrieve(id)
	if err != nil {
		return
	}

	if err = post.Delete(); err != nil {
		return
	}
	w.WriteHeader(200)
	return
}