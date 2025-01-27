package main

import (
	"fmt"

	"example.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

const balanceFile string = "balance.txt"

func main() {
	var accountBalance, err = fileops.GetFloatFromFile(balanceFile)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("------------------------")
	}

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("Reach us at", randomdata.PhoneNumber())
	fmt.Println()

	for {

		presentOptions()

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {

		case 1:
			fmt.Println("Your balance is", accountBalance)

		case 2:
			var depositAmount float64
			fmt.Print("Your deposit: ")
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}

			accountBalance += depositAmount
			fileops.WriteFloatToFile(accountBalance, balanceFile)
			fmt.Println("New balance:", accountBalance)

		case 3:
			var withdrawAmount float64
			fmt.Print("Withdraw amount: ")
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}

			if withdrawAmount > accountBalance {
				fmt.Println("Invalid amount. You can´t withdraw more than $", accountBalance, ".")
				continue
			}

			accountBalance -= withdrawAmount
			fileops.WriteFloatToFile(accountBalance, balanceFile)
			fmt.Println("New balance:", accountBalance)

		case 4:
			fmt.Println("Thanks for choosing our bank. See you later!")
			return
		}
	}
}
