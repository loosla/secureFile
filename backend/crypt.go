package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
)

// Generating a valid encryption key from the provided password.
// This key will be used to encrypt and decrypt the file content.
func createAESKey(password string) []byte {
	hash := sha256.Sum256([]byte(password))
	return hash[:16] // Use the first 16 bytes of the hash
}

// Encrypt encrypts plaintext using the given key and returns the ciphertext encoded in base64.
func encrypt(plaintext string, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// Generate a new IV (initialization vector).
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	// Encrypt the data.
	ciphertext := make([]byte, len(plaintext))
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext, []byte(plaintext))

	// Prepend the IV to the ciphertext.
	result := append(iv, ciphertext...)

	// Return the result encoded in base64.
	return base64.StdEncoding.EncodeToString(result), nil
}

// Decrypt decrypts base64-encoded ciphertext using the given key and returns the plaintext.
func decrypt(ciphertext string, key []byte) (string, error) {
	data, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	if len(data) < aes.BlockSize {
		return "", fmt.Errorf("ciphertext too short")
	}

	// Extract the IV from the beginning of the ciphertext.
	iv := data[:aes.BlockSize]
	data = data[aes.BlockSize:]

	// Decrypt the data.
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(data, data)

	return string(data), nil
}
