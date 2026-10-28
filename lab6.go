package main

import (
	"encoding/hex"
	"fmt"
)

func lab6() {
	fmt.Println("\nLab 6: XOR Encryption/Decryption")

	// Asking for user input
	plaintext := readLine("Enter plaintext: ")
	keyInput := readLine("Enter key (single character): ")
	key := keyInput[0]

	// Encrypt
	ciphertext := xorEncrypt(plaintext, key)
	fmt.Printf("Key: %q (%d)\n", key, key)
	fmt.Println("Ciphertext (hex):", hex.EncodeToString([]byte(ciphertext)))

	// Decrypt using the same function
	decrypted := xorEncrypt(ciphertext, key)
	fmt.Println("Decrypted:", decrypted)
}

func xorEncrypt(text string, key byte) string {
	result := make([]byte, len(text))
	for i := 0; i < len(text); i++ {
		result[i] = text[i] ^ key
	}
	return string(result)
}
