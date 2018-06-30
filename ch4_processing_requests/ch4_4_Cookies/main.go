package main

import (
	"net/http"
	"time"
)

func main() {
	server := &http.Server{Addr: "127.0.0.1:8080"}
	http.HandleFunc("/set_cookie", setCookie)
	server.ListenAndServe()
}

func setCookie(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{
		Name: "first_cookie",
		Value: "Go Web Programming",
		HttpOnly: true,
	}
	c2 := http.Cookie{
		Name: "second_cookide",
		Value: "Manning Publications GO",
		HttpOnly: true,
	}
	//w.Header().Set("Set-Cookie", c1.String())
	//w.Header().Add("Set-Cookie", c2.String())
	http.SetCookie(w, &c1)
	http.SetCookie(w, &c2)
}

type Cookie struct {
	Name string
	Value string
	Path string
	Domain string
	// Persistent cookie (until it expires)
	// Deprecated in favor of MaxAge
	Expires time.Time
	RawExpires string
	MaxAge int
	Secure bool
	HttpOnly bool
	Raw string
	Unparsed []string
}