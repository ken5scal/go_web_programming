package main

import (
	"net/http"
	"time"
	"fmt"
)

func main() {
	server := &http.Server{Addr: "127.0.0.1:8080"}
	http.HandleFunc("/set_cookie", setCookie)
	http.HandleFunc("/get_cookie", getCookie)
	server.ListenAndServe()
}
func getCookie(w http.ResponseWriter, r *http.Request) {
	h := r.Header["Cookie"]
	fmt.Fprint(w, "All Cookies: ")
	fmt.Fprintln(w, h)

	fmt.Fprint(w, "first_cookie: ")
	c1, err := r.Cookie("first_cookie")
	if err != nil {
		fmt.Fprintln(w, "Cannot get the first cookie")
	}
	fmt.Fprintln(w, c1)

	fmt.Fprint(w, "All Cookies: ")
	fmt.Fprintln(w, r.Cookies())
}

func setCookie(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{
		Name: "first_cookie",
		Value: "Go Web Programming",
		HttpOnly: true,
	}
	c2 := http.Cookie{
		Name: "second_cookie",
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