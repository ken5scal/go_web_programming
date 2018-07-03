package main

import (
	"os"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"io"
)

type Post struct {
	Id int `json:"id"`
	Content string `json:"content"`
	Author Author `json:"author"`
	Comments []Comment `json:"comments"`
}

type Author struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Comment struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

// Unmarshal/Decode:  parse JSON to Go Struct

// Use Decode When data is coming from an io.Reader stream,
// like the Body of an http.Request

// Use Unmarshall when data is in a string or somewhere in memory

func main() {
	jsonFile, err := os.Open("post.json")
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer jsonFile.Close()

	// Unmarshall
	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Error reading JSON data:", err)
		return
	}
	var post Post
	json.Unmarshal(jsonData, &post)
	fmt.Println(post)

	// Decode
	decoder := json.NewDecoder(jsonFile)
	for {
		var post Post
		err := decoder.Decode(&post)
		if err == io.EOF{
			break
		}
	}
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	fmt.Println(post)

	output, err := json.MarshalIndent(&post, "", "\t\t")
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}
	err = ioutil.WriteFile("post2.json", output, 0644)
	if err != nil {
		fmt.Println("Error writing JSON to file:", err)
		return
	}

	jsonFile, err = os.Create("post3.json")
	if err != nil {
		fmt.Println("Error creating JSON file:", err)
		return
	}

	if err := json.NewEncoder(jsonFile).Encode(&post); err != nil {
		fmt.Println("Error encoding JSON to file: ", err)
		return
	}
}
