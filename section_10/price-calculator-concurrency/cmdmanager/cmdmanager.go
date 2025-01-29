package cmdmanager

import "fmt"

type CmdManager struct{}

func New() CmdManager {
	return CmdManager{}
}

func (CmdManager) ReadLines() ([]string, error) {

	var prices []string
	fmt.Println("Please enter your prices. Confirm every price with ENTER")

	for {
		var price string
		fmt.Print("Price: ")
		fmt.Scan(&price)

		if price == "0" {
			break
		}

		prices = append(prices, price)
	}

	return prices, nil
}

func (CmdManager) WriteResult(data any) error {
	fmt.Println(data)
	return nil
}
