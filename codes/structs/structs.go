package main

import (
	"fmt"
)

// User is a usr type
type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
}

// Group represents a set of users
type Group struct {
	role             string
	users            []User
	newestUser       User
	isSpaceAvailable bool
}

func describeUser(user User) string {
	desc := fmt.Sprintf("Name: %s %s, Email: %s", user.FirstName, user.LastName, user.Email)
	return desc
}

func describeGroup(g Group) string {
	if len(g.users) > 2 {
		g.isSpaceAvailable = false
	}

	desc := fmt.Sprintf("Users: %d, Newest User: %s %s, Accepting New User: %t", len(g.users), g.newestUser.FirstName, g.newestUser.LastName, g.isSpaceAvailable)
	return desc
}

func main() {
	user := User{ID: 1, FirstName: "Foyez", LastName: "Ahmed", Email: "foyez@email.com"}
	user2 := User{ID: 2, FirstName: "Manam", LastName: "Ahmed", Email: "manam@email.com"}
	user3 := User{ID: 3, FirstName: "Zayan", LastName: "Ahmed", Email: "zayan@email.com"}

	group := Group{
		role:             "admin",
		users:            []User{user, user2, user3},
		newestUser:       user2,
		isSpaceAvailable: true,
	}

	fmt.Println(describeUser(user2))
	fmt.Println(describeGroup(group))
	fmt.Println(group)
}
