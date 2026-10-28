package main

import "fmt"

func lab1() {
	fmt.Println("\nLab 1: Assignment Operators")

	// Int variables
	var num1, num2 int

	//  Asking for user input
	fmt.Print("Enter two numbers (format: num1 num2): ")
	fmt.Scan(&num1, &num2)

	// Perform various assignment operations
	num1 = num2
	fmt.Println("After num1 = num2, num1:", num1)

	num1 += num2
	fmt.Println("After num1 += num2, num1:", num1)

	num1 -= num2
	fmt.Println("After num1 -= num2, num1:", num1)

	num1 *= num2
	fmt.Println("After num1 *= num2, num1:", num1)

	if num2 != 0 {
		num1 /= num2
		fmt.Println("After num1 /= num2, num1:", num1)

		num1 %= num2
		fmt.Println("After num1 %= num2, num1:", num1)
	} else {
		fmt.Print("Cannot perform division or modulo by zero.\n")
	}

}
