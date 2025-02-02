package main

type User struct {
	Id   int
	Name string
}

func (u User) Greet() string {
	return "Hello, " + u.Name
}
