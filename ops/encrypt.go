package main

import (
	"crypto/aes"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
)

func main() {
	ctb := os.Args[1]
	text := os.Args[2]
	key := getDataKey(ctb)
	encrypt(key, pad(text))
	//decrypt(key, text)
}

func encrypt(key []byte, plaintext string) {
	c, _ := aes.NewCipher(key)
	data := make([]byte, len(plaintext))
	c.Encrypt(data, []byte(plaintext))
	fmt.Println(hex.EncodeToString(data))
}

func decrypt(key []byte, ciphertext string) {
	c, _ := aes.NewCipher(key)
	cipherblob, _ := hex.DecodeString(ciphertext)
	data := make([]byte, len(cipherblob))
	c.Decrypt(data, cipherblob)
	plaintext := string(data)
	fmt.Println(unpad(plaintext))
}

func pad(text string) string {
	return text + strings.Repeat("=", 32-len(text))
}

func unpad(text string) string {
	return text[:strings.Index(text, "=")]
}

func getDataKey(cipherTextBlob string) []byte {
	region := os.Getenv("AWS_DEFAULT_REGION")
	svc := kms.New(session.New(), &aws.Config{
		Region: &region,
	})
	blob, _ := base64.StdEncoding.DecodeString(cipherTextBlob)
	result, _ := svc.Decrypt(&kms.DecryptInput{CiphertextBlob: blob})
	return result.Plaintext
}
