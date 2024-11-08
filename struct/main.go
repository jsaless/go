package main

import (
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

func newUser(firstName, lastName, birthdate string) *User {
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}
}

func (u *User) outputUserDetails() {
	fmt.Println("First Name:", u.firstName)
	fmt.Println("Last Name:", u.lastName)
	fmt.Println("Birthdate:", u.birthdate)
}

func (u *User) clearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	user := newUser(userFirstName, userLastName, userBirthdate)
	// ... do something awesome with that gathered data!

	user.outputUserDetails()
	user.clearUserName()
	user.outputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
