package main

import (
	"fmt"
	"go-struct-tag/tag"
)

type User struct {
	Id    int
	Name  string `valid:"string;maxsize(64);minsize(8)"`
	Age   int    `valid:"number;range(0,150)"`
	Email string `valid:"email"`
}

func main() {
	user := User{
		Id:    0,
		Name:  "abchello",
		Age:   0,
		Email: "hello@qq.com",
	}

	validation := tag.NewValidation(user)
	validation.Validate()
	for _, err := range validation.Errors {
		fmt.Println(err)
	}
}
