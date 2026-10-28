package main

import "fmt"

func main() {
	for {
		fmt.Println("\n===== Lab Menu =====")
		fmt.Println("1) Lab 1: Assignment Operators")
		fmt.Println("2) Lab 2: Logical Operators")
		fmt.Println("3) Lab 3: Bitwise and Assignment Operators")
		fmt.Println("4) Lab 4: Mini Calculator")
		fmt.Println("5) Lab 5: Binary, Hex, and Base64 Encoding")
		fmt.Println("6) Lab 6: XOR Encryption/Decryption")
		fmt.Println("0) Exit")

		var choice int
		fmt.Print("Choose: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			lab1()
		case 2:
			lab2()
		case 3:
			lab3()
		case 4:
			lab4()
		case 5:
			lab5()
		case 6:
			lab6()
		case 0:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}
