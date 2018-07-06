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
	mux.HandleFunc("/post/", handleRequest)
	writer = httptest.NewRecorder()    //captures returne http response
}

func TestHandleGet(t *testing.T)  {
	posdId := 2
	request, _:= http.NewRequest("GET", fmt.Sprintf("/post/%d", posdId), nil)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}

	var post Post
	json.Unmarshal(writer.Body.Bytes(), &post)
	if post.Id != posdId {
		t.Errorf("Cannot retrieve JSON post")
	}
}

func TestHandlePut(t *testing.T) {
	json  := strings.NewReader(`{"content":"Updated post","author":"Kengo Suzuki"}`)
	request, _ := http.NewRequest("PUT", "/post/1", json)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
}