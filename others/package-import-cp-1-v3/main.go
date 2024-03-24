package main

import (
	"a21hc3NpZ25tZW50/internal"
	"fmt"
	"strconv"
	"strings"
)

func AdvanceCalculator(calculate string) float32 {
	parts := strings.Split(calculate, " ")
	if len(parts) == 0 {
		return 0
	}
	number, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0
	}
	fmt.Println(parts)

	calc := internal.NewCalculator(float32(number))
	for i := 1; i < len(parts); i += 2 {
		operator := parts[i]
		operand, err := strconv.ParseFloat(parts[i+1], 32)
		if err != nil {
			return 0
		}
		switch operator {
		case "+":
			calc.Add(float32(operand))
		case "-":
			calc.Subtract(float32(operand))
		case "*":
			calc.Multiply(float32(operand))
		case "/":
			calc.Divide(float32(operand))
		default:
			return 0
		}
	}
	return calc.Result() // TODO: replace this
}

func main() {
	res := AdvanceCalculator("3 * 4 / 2 + 10 - 5")
	res2 := AdvanceCalculator("10 + 2")

	fmt.Println(res)
	fmt.Println(res2)
}
