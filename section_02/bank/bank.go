package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const balanceFile string = "balance.txt"

func getBalanceFromFile() (float64, error) {
	var data, readErr = os.ReadFile(balanceFile)

	if readErr != nil {
		return 1000, errors.New("Failed to read balance file.")
	}

	var balanceText = string(data)
	var balance, parseErr = strconv.ParseFloat(balanceText, 64)

	if parseErr != nil {
		return 1000, errors.New("Failed to parse stored balance value.")
	}

	return balance, nil
}

func writeBalanceToFile(balance float64) {
	var balanceText = fmt.Sprint(balance)
	os.WriteFile(balanceFile, []byte(balanceText), 0644)
}

func main() {
	var accountBalance, err = getBalanceFromFile()

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("------------------------")
		panic("Can't continue, sorry...")
	}

	fmt.Println("Welcome to Go Bank!")

	for {

		fmt.Println("What do you want to do?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

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
			writeBalanceToFile(accountBalance)
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
			writeBalanceToFile(accountBalance)
			fmt.Println("New balance:", accountBalance)

		case 4:
			fmt.Println("Thanks for choosing our bank. See you later!")
			return
		}
	}
}
