package main

import (
	"fmt"
)

type User struct {
	Name  string
	Email string
}

type Admin struct {
	User
	Level int
}

func main() {
	u := User{
		Name:  "Bob",
		Email: "Bob@email.com",
	}

	fmt.Printf("User = %v\n", u)

	a := Admin{
		Level: 5,
		User: User{
			Name: "Aurel",
			Email: "aurel@gmail.com",
		},
	}

	admin := Admin{
		Level: 2,
	}

	admin.Name = "Alice"
	admin.Email = "alice@org.com"
	fmt.Printf("Admin = %v\n", admin)

	fmt.Println("Declaration différente")
	fmt.Printf("Admin = %v\n", a)
}
