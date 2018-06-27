package main

import (
	"net/http"
	"fmt"
	"runtime"
	"reflect"
)

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
		//Handler: &MyHandler{}, // This gonna route everything here
	}

	//type Handler interface {
	//  ServeHTTP(ResponseWriter, *Request)
	//}

	//type HandlerFunc func(ResponseWriter, *Request)
	//func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	//  f(w, r)
	//}

	// http.HandleFunc("/hello", log(hello)) // This works as well
	http.Handle("/hello", log(hello))  //<- Handlerfunc implements Handler so this works too
	http.Handle("/protect", protect(log(hello)))
	//http.HandleFunc("/protect",protect(log(hello))) //<- This won't work because Handler cannot be HandlerFunc

	// For any registered URLs that don’t end with a slash (/),
	// ServeMux will try to match the exact URL pattern.
	// If the URL ends with a slash (/), ServeMux will see
	// if the requested URL starts with any registered URL
	http.Handle("/hello/", log(hello))

	server.ListenAndServe()
}

type MyHandler struct{}
func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the default!")
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

func log(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Println("Handler function called - " + name)
		h(w, r)
	}
}

func protect(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	})
}