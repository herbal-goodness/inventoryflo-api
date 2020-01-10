package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
)

func main() {
	ctb := os.Args[1]
	text := os.Args[2]
	c, _ := aes.NewCipher(getDataKey(ctb))
	gcm, _ := cipher.NewGCM(c)
	encrypt(gcm, text)
	//decrypt(gcm, text)
}

func encrypt(gcm cipher.AEAD, plaintext string) {
	nonce := make([]byte, gcm.NonceSize())
	io.ReadFull(rand.Reader, nonce)
	data := gcm.Seal(nonce, nonce, []byte(plaintext), nil)
	fmt.Println(hex.EncodeToString(data))
}

func decrypt(gcm cipher.AEAD, cipherText string) {
	cipherBlob, _ := hex.DecodeString(cipherText)
	ns := gcm.NonceSize()
	nonce, blob := cipherBlob[:ns], cipherBlob[ns:]
	plain, _ := gcm.Open(nil, nonce, blob, nil)
	fmt.Println(string(plain))
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
