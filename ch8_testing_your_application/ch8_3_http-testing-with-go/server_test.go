package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"encoding/json"
)

func TestHandleGet(t *testing.T)  {
	mux := http.NewServeMux()
	mux.HandleFunc("/post/", handleRequest)

	writer := httptest.NewRecorder()    //captures returne http response
	request, _:= http.NewRequest("GET", "/post/8", nil)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}

	var post Post
	json.Unmarshal(writer.Body.Bytes(), &post)
	if post.Id != 1 {
		t.Error("Cannot retrieve JSON post")
	}
}