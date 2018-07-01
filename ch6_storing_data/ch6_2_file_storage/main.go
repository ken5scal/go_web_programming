package main

import (
	"io/ioutil"
	"fmt"
	"os"
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
}
