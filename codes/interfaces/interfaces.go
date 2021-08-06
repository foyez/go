package main

import (
	"fmt"
)

// Describer has a Describe method
type Human interface {
	Describe() string
}

// Any struct that has a method called describe is also a type of Describer

type User struct {
	ID                         int
	FirstName, LastName, Email string
}

type Group struct {
	role           string
	users          []User
	newestUser     User
	spaceAvailable bool
}

func (u User) Describe() string {
	desc := fmt.Sprintf("Name: %s %s, Email: %s", u.FirstName, u.LastName, u.Email)
	return desc
}

func (g *Group) Describe() string {
	if len(g.users) > 2 {
		g.spaceAvailable = false
	}

	desc := fmt.Sprintf("Users: %d, Newest User: %s %s, Accepting New User: %t", len(g.users), g.newestUser.FirstName, g.newestUser.LastName, g.spaceAvailable)
	return desc
}

func main() {
	user := User{ID: 1, FirstName: "Foyez", LastName: "Ahmed", Email: "foyez@email.com"}
	user2 := User{ID: 2, FirstName: "Manam", LastName: "Ahmed", Email: "manam@email.com"}

	group := Group{
		role:           "admin",
		users:          []User{user, user2},
		newestUser:     user2,
		spaceAvailable: true,
	}

	describeSomething := func(human Human) {
		desc := human.Describe()
		fmt.Println(desc)
	}

	describeSomething(user)
	describeSomething(&group)
}
