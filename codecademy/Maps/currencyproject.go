package main

import (
	"fmt"
)

func main() {
	currencies := map[string]float32{
		"GBP": 0.79,
		"EUR": 0.93,
		"AED": 3.67,
	}

	var dollarAmount float32
	var currency string

	fmt.Println("Enter an amount in Dollars")
	fmt.Scan(&dollarAmount)

	if dollarAmount == 0 {
		fmt.Println("Input is not a valid float. Try again")
	} else {
		fmt.Println("Enter the currency you want: E.g. GBP")
		fmt.Scan(&currency)

		rate, isValid := currencies[currency]

		if isValid {
			targetCurrency := rate * dollarAmount
			fmt.Println("Your new currency amount is", targetCurrency, currency)
		} else {
			fmt.Println("Currency not available")
		}
	}
}

