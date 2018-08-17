package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"encoding/json"
	"strings"
	"os"
	"fmt"
)

var mux *http.ServeMux
var writer *httptest.ResponseRecorder

func TestMain(m *testing.M) {
	setUp()           // runs only once for all test cases
	code := m.Run()  //  The individual test case functions are called
	os.Exit(code)
}

func setUp() {
	mux = http.NewServeMux()
	mux.HandleFunc("/post/", handleRequest(&FakePost{}))
	writer = httptest.NewRecorder()    //captures returne http response
}

func TestHandleGet(t *testing.T)  {
	postId := 2
	request, _:= http.NewRequest("GET", fmt.Sprintf("/post/%d", postId), nil)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}

	var post Post
	json.Unmarshal(writer.Body.Bytes(), &post)
	if post.Id != postId {
		t.Errorf("Cannot retrieve JSON post")
	}
}

func TestHandlePut(t *testing.T) {
	postId := 2
	json  := strings.NewReader(`{"content":"Updated post","author":"Kengo Suzuki"}`)
	request, _ := http.NewRequest("PUT", fmt.Sprintf("/post/%d", postId), json)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
}