package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
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
	//r.ParseForm()
	//fmt.Fprint(w, "r.Form: ")
	//fmt.Fprintln(w,  r.Form)
	//
	//// only get the value to the key `post`, form value is always prioritized
	//fmt.Fprint(w, "r.Form[\"post\": ")
	//fmt.Fprintln(w, r.Form["post"])
	//
	//// just  get the form key-value pairs. Only supports application/x-www-form-urlencoded
	//fmt.Fprint(w, "r.PostForm: ")
	//fmt.Fprintln(w, r.PostForm)

	// Parsing MultipartForm Encode data
	//r.ParseMultipartForm(1024)
	//fileHeader := r.MultipartForm.File["uploaded"][0]
	//file, err := fileHeader.Open()
	// FormFile returns the first value given the key
	// Good if you have only one file to be uploaded
	file, _, err := r.FormFile("uploaded")
	if err == nil {
		data, err := ioutil.ReadAll(file)
		if err == nil {
			fmt.Fprintln(w, string(data))
		}
	}
	//
	//// return the form key-value pairs in struct form
	//fmt.Fprint(w, "r.MultipartForm: ")
	//fmt.Fprintln(w, r.MultipartForm)

	// Use this to get the key-value pairs for enctype application/x-www-form urlencoded
	// no need to call ParseFOrm/ParseMaltiPartForm
	//fmt.Fprint(w, "r.FormValue: ")
	//fmt.Fprintln(w, r.FormValue("hello"))
	//
	//// PostFormValue call the ParseMaultipart method
	//// but if enctype is `multipart/form-data, then it won't work
	//fmt.Fprint(w, "r.PostFormValue(\"hello\"): ")
	//fmt.Fprintln(w, r.PostFormValue("hello"))
	//fmt.Fprint(w, "r.PostForm: ")
	//fmt.Fprintln(w, r.PostForm) // Must call `Form` method first
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

