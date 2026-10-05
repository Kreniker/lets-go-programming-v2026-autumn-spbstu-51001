package main

import (
	"fmt"
)

func processRequest(operator string, num int, minNum *int, maxNum *int) int {
	if operator == ">=" {
		if num > *minNum {
			*minNum = num
		}
	} else if operator == "<=" {
		if num < *maxNum {
			*maxNum = num
		}
	}

	if *minNum > *maxNum {
		return -1
	}

	return *minNum
}

func main() {
	var officesCount int
	_, err := fmt.Scan(&officesCount)
	if err != nil {
		fmt.Println("Non-correct count offices")

		return
	}
	var i int = 0
	for ; i < officesCount; i++ {
		var employeeCount int

		_, err := fmt.Scan(&employeeCount)
		if err != nil {
			return
		}
		var minNum int = 15
		var maxNum int = 30
		var j int = 0
		needLen := 2
		for ; j < employeeCount; j++ {
			var operator string
			_, err := fmt.Scan(&operator)
			if err != nil {
				return
			}

			if len(operator) != needLen {
				return
			}

			if operator != "<=" && operator != ">=" {
				return
			}

			var num int
			_, err = fmt.Scan(&num)

			if err != nil {
				return
			}

			ans := processRequest(operator, num, &minNum, &maxNum)

			fmt.Println(ans)
		}
	}
}
