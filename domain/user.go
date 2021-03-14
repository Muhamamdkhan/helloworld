package domain

import "fmt"

type User struct {
	Firstname string
	Lastname  string
	Email     string
}

func (u User) SayHello() {
	fmt.Println("Hello, ", u.Firstname)
}
