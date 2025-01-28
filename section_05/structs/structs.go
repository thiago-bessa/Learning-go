package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (DD/MM/YYYY): ")

	var appUser, err = user.New(firstName, lastName, birthdate)

	if err != nil {
		fmt.Println(err)
		return
	}

	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()

	admin := user.NewAdmin("test@example.com", "test123")
	admin.OutputUserDetails()
	admin.ClearUserName()
	admin.OutputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
