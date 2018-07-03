package main

import (
	_ "github.com/lib/pq"
	"github.com/jinzhu/gorm"
	"time"
	"fmt"
)

type Post struct {
	Id int
	Content string
	Author string `sql:"not null"`
	Comments []Comment
	CreatedAt time.Time
}

type Comment struct {
	Id int
	Content string
	Author string `sql:"not null"`
	PostId int `sql:"index"`
	CreatedAt time.Time
}

var Db *gorm.DB

func init() {
	var err error
	Db, err = gorm.Open("postgres", "user=gwp dbname=gwp password=gwp sslmode=disable")
	if err != nil {
		panic(err)
	}
	Db.AutoMigrate(&Post{}, &Comment{})
}

func main() {
	post := Post{Content: "Hello World!", Author: "Sau Sheong"}
	fmt.Println(post)

	Db.Create(&post)
	fmt.Println(post)

	comment := Comment{Content: "Good post!", Author: "Joe"}
	Db.Model(&post).Association("Comments").Append(comment)

	var readPost Post
	Db.Where("author = $1", "Sau Sheong").First(&readPost)
	var comments []Comment
	Db.Model(&readPost).Related(&comments)
	fmt.Println(comments[0])
}
