package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"

	"golang.org/x/crypto/pbkdf2"
)

var sessionKey []byte

func deriveKey(password, salt string) []byte {
	saltBytes, err := base64.StdEncoding.DecodeString(salt)
	if err != nil {
		fmt.Println("Error decoding salt:", err)
		return nil
	}
	key := pbkdf2.Key([]byte(password), saltBytes, 4096, 32, sha256.New)
	return key
}

func encrypt(text string, key []byte) string {
	block, _ := aes.NewCipher(key)
	gcm, _ := cipher.NewGCM(block)

	nonce := make([]byte, gcm.NonceSize())
	io.ReadFull(rand.Reader, nonce)

	cipherText := gcm.Seal(nonce, nonce, []byte(text), nil)
	return base64.StdEncoding.EncodeToString(cipherText)
}

func decrypt(cipherText string, key []byte) string {
	data, _ := base64.StdEncoding.DecodeString(cipherText)

	block, _ := aes.NewCipher(key)
	gcm, _ := cipher.NewGCM(block)

	nonceSize := gcm.NonceSize()
	nonce := data[:nonceSize]
	cipherTextData := data[nonceSize:]

	text, _ := gcm.Open(nil, nonce, cipherTextData, nil)
	return string(text)
}

func generateSalt() string {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(salt)
}
