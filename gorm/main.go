package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

type User struct {
	ID   int64
	UID  int
	Name string
}

func Select() {
	first := &User{}
	db.Table("user").First(first)
	fmt.Println(first)

	last := &User{}
	db.Table("user").Last(last)
	fmt.Println(last)

	user := &User{}
	db.Table("user").Where(map[string]interface{}{"ID": 2}).Find(user)
	fmt.Println(user)
}

func Update() {
	user := &User{
		ID:   1,
		UID:  1,
		Name: "新名字",
	}
	db.Table("user").Save(user)

	nuser := &User{}
	db.Table("user").First(nuser)
	fmt.Println(nuser)
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/test"
	db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	Select()
	Update()
}
