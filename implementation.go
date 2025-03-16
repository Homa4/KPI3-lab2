package lab2

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// TODO: document this function.
// PrefixToPostfix converts
func CalculatePrefix(expression string) (int, error) {
	tokens := strings.Fields(expression)
	var stack []int

	for i := len(tokens) - 1; i >= 0; i-- {
		token := tokens[i]

		switch token {
		case "+", "-", "*", "/", "^":
			if len(stack) < 2 {
				return 0, fmt.Errorf("invalid expression")
			}
			a, b := stack[len(stack)-1], stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			var res int
			switch token {

			case "^":
				res = int(math.Pow(float64(a), float64(b)))
			case "+":
				res = a + b
			case "-":
				res = a - b
			case "*":
				res = a * b
			case "/":
				if b == 0 {
					return 0, fmt.Errorf("division by zero")
				}
				res = a / b
			}
			stack = append(stack, res)
		default:
			num, err := strconv.Atoi(token)
			if err != nil {
				return 0, fmt.Errorf("invalid token: %s", token)
			}
			stack = append(stack, num)
			fmt.Print(stack)
		}
	}

	return stack[0], nil
}
