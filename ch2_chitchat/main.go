package main

import (
	"net/http"
	"html/template"
	"github.com/sausheong/gwp/Chapter_2_Go_ChitChat/chitchat/data"
)

func main() {
	files := http.FileServer(http.Dir("/public"))

	// multiplexer: redirects a request to a handler.
	mux := http.NewServeMux()
	mux.Handle("/static", http.StripPrefix("/static", files))
	mux.HandleFunc("/", index)
	mux.HandleFunc("/err", err)

	mux.HandleFunc("/login", login)
	mux.HandleFunc("/logout", logout)
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/signup_account", signupAccount)
	mux.HandleFunc("/authenticate", authenticate)
	mux.HandleFunc("/thread/new", newThread)
	mux.HandleFunc("/thread/create", createThread)
	mux.HandleFunc("/thread/post", postThread)
	mux.HandleFunc("/thread/read", readThread)

	server := &http.Server{
		Addr: "0.0.0.0/8080",
		Handler: mux,
	}
	server.ListenAndServe()
}

func index(writer http.ResponseWriter, request *http.Request) {
	threads, err := data.Threads(); if err == nil {
		_, err := session(writer, request)
		public_tmpl_filesfiles := []string{
			"templates/layout.html",
			"templates/public.navbar.html",
			"templates/index.html"}
		private_tmpl_filesfiles := []string{
			"templates/layout.html",
			"templates/private.navbar.html",
			"templates/index.html"}
		var templates *template.Template
		if err != nil {
			templates := template.Must(template.ParseFiles(private_tmpl_filesfiles...))
		} else {
			templates := template.Must(template.ParseFiles(public_tmpl_filesfiles...))
		}
		templates.ExecuteTemplate(w, "layout", threads)
	}
}

