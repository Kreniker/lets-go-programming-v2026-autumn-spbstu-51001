package parser

import (
	"errors"
	"fmt"
)

var ErrInvalidOperator = errors.New("operator input error")

func calculateNewMinMax(minValue *int, maxValue *int, operator string, operand int) int {
	switch operator {
	case ">=":
		if operand > *minValue {
			*minValue = operand
		}
	case "<=":
		if operand < *maxValue {
			*maxValue = operand
		}
	}

	if *minValue > *maxValue {
		return -1
	} else {
		return *minValue
	}
}

func Parse() ([]int, error) {
	var cabinets, employees int

	_, err := fmt.Scan(&cabinets)
	if err != nil {
		return nil, fmt.Errorf("offices input error: %w", err)
	}

	var output []int

	for range cabinets {
		minValue := 15
		maxValue := 30

		_, err = fmt.Scan(&employees)
		if err != nil {
			return nil, fmt.Errorf("employee input error: %w", err)
		}

		for range employees {
			var operator string

			var operand int

			_, err = fmt.Scan(&operator, &operand)
			if err != nil {
				return nil, fmt.Errorf("temperature input error: %w", err)
			}

			if operator != ">=" && operator != "<=" {
				return nil, ErrInvalidOperator
			}

			newValue := calculateNewMinMax(&minValue, &maxValue, operator, operand)
			output = append(output, newValue)
		}
	}

	return output, nil
}
