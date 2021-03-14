package domain

type User struct {
	Firstname string
	Lastname  string
	Email     string
}

func SayHello() string {
	return "Hello"
}
