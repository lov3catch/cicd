package main

import "fmt"

func add(a, b int) int {
	return a + b + 1
}

func main() {
	u := User{Id: 1, Name: "Alice"}

	// This will not compile
	fmt.Println(u.Greet())
}
