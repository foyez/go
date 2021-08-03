package main

import "fmt"

func main() {
	var results map[string]float64 = make(map[string]float64)

	results["foyez"] = 3.4
	results["mithu"] = 3.5

	fmt.Println(results) // map[foyez:3.4 mithu:3.5]

	// ***********************************************
	userEmails := map[int]string{
		1: "user1@email.com",
		2: "user2@email.com",
	}

	userEmails[1] = "user12@email.com"
	emailOfSecondUser, ok := userEmails[2]
	emailOfFourthUser, ok2 := userEmails[4]

	fmt.Println(userEmails)             // map[1:user12@email.com 2:user2@email.com]
	fmt.Println(emailOfSecondUser, ok)  // user2@email.com true
	fmt.Println(emailOfFourthUser, ok2) // false

	if email, ok := userEmails[2]; ok {
		fmt.Printf("%s exists\n", email)
	} else {
		fmt.Printf("%s doesn't exists\n", email)
	}

	delete(userEmails, 1)
	fmt.Println(userEmails) // [2:user2@email.com]
}
