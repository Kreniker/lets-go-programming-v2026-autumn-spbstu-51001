package main

import (
	"fmt"
)

func main() {
	var first_argument int
	_, err := fmt.Scan(&first_argument)
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}
	var second_argument int
	_, err = fmt.Scan(&second_argument)
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
}