package main

import (
	"fmt"
)

func main() {
	var firstArgument int
	_, err := fmt.Scan(&firstArgument)
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}
	var secondArgument int
	_, err = fmt.Scan(&secondArgument)
	if err != nil {
		fmt.Println("Invalid second operand")
		return
	}
	var operation string
	_, err = fmt.Scan(&operation)
	if err != nil {
		fmt.Println("Invalid operation")
		return
	}
	switch operation {
	case "+":
		fmt.Println(firstArgument + secondArgument)
	case "-":
		fmt.Println(firstArgument - secondArgument)
	case "*":
		fmt.Println(firstArgument - secondArgument)
	case "/":
		if secondArgument == 0 {
			fmt.Println("Division by zero")
			break
		}
		fmt.Println(firstArgument / secondArgument)
	default:
		fmt.Println()
	}
}
