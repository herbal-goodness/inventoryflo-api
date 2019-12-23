package main

import (
	"crypto/aes"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

func main()  {
	key := []byte(os.Args[1])
	text := os.Args[2]
	encrypt(key, pad(text))
	//decrypt(key, os.Args[2])
}

func encrypt(key []byte, plaintext string) {
	c, _ := aes.NewCipher(key[:32])
	data := make([]byte, len(plaintext))
	c.Encrypt(data, []byte(plaintext))
	fmt.Println(hex.EncodeToString(data))
}

func decrypt(key []byte, ciphertext string) {
	c, _ := aes.NewCipher(key[:32])
	cipherblob, _ := hex.DecodeString(ciphertext)
	data := make([]byte, len(cipherblob))
	c.Decrypt(data, cipherblob)
	fmt.Println(unpad(string(data[:])))
}

func pad(text string) string {
	return text + strings.Repeat("^", 32-len(text))
}

func unpad(text string) string {
	return strings.ReplaceAll(text, "^", "")
}