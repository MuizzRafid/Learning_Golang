package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	u := User{Name: "july", Age: 6}
	fmt.Println("before",u.Age)
	u.Birthday()
	fmt.Println("after",u.Age)
}

func (u *User) Birthday() {
	u.Age++
}