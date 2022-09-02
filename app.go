package main

type User struct {
	Id    int
	Name  string `valid:"string;maxsize(64);minsize(8)"`
	Age   int    `valid:"number;range(0,150)"`
	Email string `valid:"email"`
}

func main() {

}
