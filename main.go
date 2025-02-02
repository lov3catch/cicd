package main

import "fmt"

func main() {
	u := User{Id: 1, Name: "Alice"}

	// This will not compile
	fmt.Println(u.Greet())
}
