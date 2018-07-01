package main

import (
	"io/ioutil"
	"fmt"
	"os"
	"encoding/csv"
	"strconv"
)

func main() {
	data := []byte("hello world!\n")
	if err := ioutil.WriteFile("data1", data, 0644); err != nil {
		panic(err)
	}
	read1, _ := ioutil.ReadFile("data1")
	fmt.Print(string(read1))

	file1, _:= os.Create("data2")
	defer file1.Close()
	file1.Write(data)

	file2, _ := os.Open("data2")
	defer file2.Close()
	b := make([]byte, len(data))
	file2.Read(b)
	fmt.Println(string(b))

	// Create CSV
	csvFile, err := os.Create("posts.csv")
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	allPosts := []Post{
		Post{Id: 1, Content: "Hello World!", Author: "Kengo Suzuki"},
	}

	// Write to CSV
	writer := csv.NewWriter(csvFile)
	for _, post := range allPosts {
		line := []string{strconv.Itoa(post.Id), post.Content, post.Author}
		if err := writer.Write(line); err != nil {
			panic(err)
		}
	}
	// Make sure that any buffered data is properly written to the file
	writer.Flush()

	// Read from CSV
	file, err := os.Open("posts.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	// set positive number, if the number of fields you expect
	// from each record and Go will throw an error
	// if you get less from the CSV file
	reader.FieldsPerRecord = -1
	record, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	var posts []Post
	for _, item := range record {
		id, _ := strconv.ParseInt(item[0], 0, 0)
		post := Post{Id: int(id), Content: item[1], Author: item[2]}
		posts = append(posts, post)
	}

	fmt.Println(posts[0].Id)
	fmt.Println(posts[0].Content)
	fmt.Println(posts[0].Author)
}

type Post struct {
	Id int
	Content string
	Author string
}