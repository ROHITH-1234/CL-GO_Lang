package main

import (
	"fmt"
	"os"
)

func switchstatement() {
	var input int
	fmt.Println("Enter the value:")
	fmt.Fscan(os.Stdin, &input)
	switch input {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Satday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid data")
	}
}
