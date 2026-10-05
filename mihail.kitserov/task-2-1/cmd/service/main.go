package main

import (
	"fmt"
)

func main() {
	var offices_count uint
	_, err := fmt.Scan(&offices_count)
	if err != nil {
		fmt.Println("Non-correct count offices")
		return
	}
	for i := 0; i < int(offices_count); i++ {
		var employee_count uint
		_, err := fmt.Scan(&employee_count)
		if err != nil {
			fmt.Println("Non-correct employee count")
			return
		}
		min := 15
		max := 30
		for j := 0; j < int(employee_count); j++ {
			var operator string
			_, err := fmt.Scan(&operator)
			if err != nil {
				fmt.Println("Non-correct operator")
				return
			}
			if len(operator) != 2 {
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
				max = int(num)
			} else {
				min = int(num)
			}
			if min > max {
				fmt.Println("-1")
				break
			}
			fmt.Println(min)
		}
	}
}
