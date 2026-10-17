package main

import (
	"fmt"

	"github.com/electricbayan/task-2-1/internal/parser"
)

func main() {
	var output []int

	output, err := parser.Parse()
	if err != nil {
		return
	}

	for _, num := range output {
		fmt.Println(num)
	}
}
