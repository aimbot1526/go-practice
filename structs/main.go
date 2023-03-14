package main

import "fmt"

func main() {
	newUser := User{"test user", "myemail@sm.com", 24}
	fmt.Println(newUser)
}

type User struct {
	Name  string
	Email string
	age   int
}
