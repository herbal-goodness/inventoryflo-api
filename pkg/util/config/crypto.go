package config

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
)

var blockCipher cipher.Block = nil

func decrypt(ciphertext string) string {
	if blockCipher == nil {
		blockCipher, _ = aes.NewCipher(getDataKey())
	}
	cipherblob, _ := hex.DecodeString(ciphertext)
	data := make([]byte, len(cipherblob))
	blockCipher.Decrypt(data, cipherblob)
	return unpad(string(data))
}

func getDataKey() []byte {
	svc := kms.New(session.New())
	blob, _ := base64.StdEncoding.DecodeString(os.Getenv("data_key"))
	result, _ := svc.Decrypt(&kms.DecryptInput{CiphertextBlob: blob})
	return result.Plaintext
}

func unpad(text string) string {
	return text[:strings.Index(text, "=")]
}
