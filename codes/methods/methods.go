package main

import (
	"fmt"
)

type User struct {
	ID                         int
	FirstName, LastName, Email string
}

func describeUser(u User) string {
	desc := fmt.Sprintf("Name: %s %s, Email: %s, ID: %d", u.FirstName, u.LastName, u.Email, u.ID)
	return desc
}

// func (receiverName ReceiverType) MethodName(args)
func (u *User) Describe() string {
	desc := fmt.Sprintf("Name: %s %s, Email: %s, ID: %d", u.FirstName, u.LastName, u.Email, u.ID)
	return desc
}

func main() {
	user := User{ID: 1, FirstName: "Manam", LastName: "Ahmed", Email: "chayon@email.com"}

	desc := describeUser(user)
	desc2 := user.Describe()

	fmt.Println(desc)
	fmt.Println(desc2)
}
