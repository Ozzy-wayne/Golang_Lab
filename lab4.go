package main

import (
	"errors"
	"fmt"
)

func lab4() {
	for {
		fmt.Println("\n===== Mini Calculator =====")
		fmt.Println("1) Add  2) Sub  3) Mul  4) Div  5) Mod  6) Exit")

		var choice int
		fmt.Print("Choose: ")
		fmt.Scan(&choice)

		if choice == 6 {
			fmt.Println("Exiting calculator...")
			return
		}
		if choice < 1 || choice > 6 {
			fmt.Println("Error: invalid operator")
			continue
		}

		// Int variables
		var a, b int
		fmt.Print("Enter a: ")
		fmt.Scan(&a)
		fmt.Print("Enter b: ")
		fmt.Scan(&b)

		result, err := calculate(a, b, choice)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		fmt.Println("Result:", result)
	}
}

func calculate(a, b, op int) (int, error) {
	switch op {
	case 1:
		return add(a, b), nil
	case 2:
		return sub(a, b), nil
	case 3:
		return mul(a, b), nil
	case 4:
		return div(a, b)
	case 5:
		return mod(a, b)
	default:
		return 0, errors.New("invalid operator")
	}
}

func add(a, b int) int { return a + b }
func sub(a, b int) int { return a - b }
func mul(a, b int) int { return a * b }

func div(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func mod(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a % b, nil
}
