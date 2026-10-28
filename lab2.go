package main

import "fmt"

func lab2() {
	fmt.Println("\nLab 2: Logical Operators")

	// Int variables
	var num1, num2 int

	//  Asking for user input
	fmt.Print("Enter two numbers (format: num1 num2): ")
	fmt.Scan(&num1, &num2)

	// Perform Logical Expressions
	result := (num1 > 0) && (num2 > 0)
	fmt.Printf("%d && %d (both positive): %t\n", num1, num2, result)

	result = (num1 > num2) || (num2 > num1)
	fmt.Printf("%d || %d (One larger than another): %t\n", num1, num2, result)

	result = !(num1 == num2)
	fmt.Printf("!(%d == %d) (not equal): %t\n", num1, num2, result)
}
