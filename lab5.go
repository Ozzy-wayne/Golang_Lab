package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

func lab5() {
	fmt.Println("\nLab 5: Binary, Hex, and Base64 Encoding")

	// Asking for user input (read full line so spaces are kept)
	text := readLine("Enter a string: ")

	// Binary representation (8 bits per byte)
	var binary []string
	for _, b := range []byte(text) {
		binary = append(binary, fmt.Sprintf("%08b", b))
	}
	fmt.Println("Binary:", strings.Join(binary, " "))

	// Hexadecimal representation
	fmt.Println("Hex:   ", hex.EncodeToString([]byte(text)))

	// Base64 representation
	fmt.Println("Base64:", base64.StdEncoding.EncodeToString([]byte(text)))
}

// readLine reads a whole line (keeps spaces), skipping the leftover newline from fmt.Scan
func readLine(prompt string) string {
	fmt.Print(prompt)
	for {
		var line []byte
		buf := make([]byte, 1)
		for {
			n, err := os.Stdin.Read(buf)
			if n == 0 || err != nil || buf[0] == '\n' {
				break
			}
			line = append(line, buf[0])
		}
		text := strings.TrimRight(string(line), "\r")
		if text != "" {
			return text
		}
	}
}
