package main

import (
	"fmt"
)

func main() {
	var officesCount uint
	_, err := fmt.Scan(&officesCount)
	if err != nil {
		fmt.Println("Non-correct count offices")

		return
	}
	var i uint = 0
	for ; i < officesCount; i++ {
		var employeeCount uint
		_, err := fmt.Scan(&employeeCount)
		if err != nil {
			fmt.Println("Non-correct employee count")

			return
		}
		minNum := 15
		maxNum := 30
		var j uint = 0
		needLen := 2
		for ; j < employeeCount; j++ {
			var operator string
			_, err := fmt.Scan(&operator)
			if err != nil {
				fmt.Println("Non-correct operator")

				return
			}

			if len(operator) != needLen {
				fmt.Println("Too long or short operator")

				return
			}

			if operator != "<=" && operator != ">=" {
				fmt.Println("First need to be '<=' or '>='")

				return
			}

			var num uint
			_, err = fmt.Scan(&num)

			if err != nil {
				fmt.Println("Non-correct number")

				return
			}

			if num > 30 || num < 15 {
				fmt.Println("Number must be in range [15, 30]")

				return
			}

			if operator == "<=" {
				maxNum = int(num)
			} else {
				minNum = int(num)
			}

			if minNum > maxNum {
				fmt.Println("-1")

				continue
			}
			
			fmt.Println(minNum)
		}
	}
}
