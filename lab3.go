package main

import "fmt"

// Bitwise Operations Functions
func myAND(a, b int) {
	fmt.Printf("AND of %04b and %04b is %04b\n", a, b, a&b)
}

func myOR(a, b int) {
	fmt.Printf("OR of %04b and %04b is %04b\n", a, b, a|b)
}

func myXOR(a, b int) {
	fmt.Printf("XOR of %04b and %04b is %04b\n", a, b, a^b)
}

func myNOT(a, b int) {
	fmt.Printf("NOT of %04b is %04b\n", a, ^a&0b1111)
	fmt.Printf("NOT of %04b is %04b\n", b, ^b&0b1111)
}

func myShift(a int, n uint) {
	fmt.Printf("Left Shift of %04b by %d is %04b\n", a, n, a<<n)
	fmt.Printf("Right Shift of %04b by %d is %04b\n", a, n, a>>n)
}

// Lab3 function to demonstrate bitwise operations
func lab3() {
	fmt.Println("\nLab3: Bitwise Operators")

	var a, b int
	fmt.Print("Enter two integers (format: num1 num2): ")
	fmt.Scan(&a, &b)
	fmt.Printf("%d = %04b\n", a, a)
	fmt.Printf("%d = %04b\n", b, b)

	fmt.Println("\n--Bitwise Operations--")

	myAND(a, b)
	myOR(a, b)
	myXOR(a, b)
	myNOT(a, b)
	myShift(a, 2)

}
