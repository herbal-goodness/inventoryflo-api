package config

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
)

var gcm cipher.AEAD = nil

func decrypt(cipherText string) string {
	if gcm == nil {
		c, _ := aes.NewCipher(getDataKey())
		gcm, _ = cipher.NewGCM(c)
	}
	cipherBlob, _ := hex.DecodeString(cipherText)
	ns := gcm.NonceSize()
	nonce, blob := cipherBlob[:ns], cipherBlob[ns:]
	plain, _ := gcm.Open(nil, nonce, blob, nil)
	return string(plain)
}

func getDataKey() []byte {
	svc := kms.New(session.New())
	blob, _ := base64.StdEncoding.DecodeString(os.Getenv("data_key"))
	result, _ := svc.Decrypt(&kms.DecryptInput{CiphertextBlob: blob})
	return result.Plaintext
}
