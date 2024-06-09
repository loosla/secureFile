package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
)

// Generate a new AES key of the specified size in bytes (e.g., 32 for AES-256).
func generateKey(size int) ([]byte, error) {
	key := make([]byte, size)
	if _, err := rand.Read(key); err != nil {
		return nil, err
	}

	return key, nil
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

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
}

func main() {

	http.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		if r.Method == http.MethodOptions {
			return
		}
		fmt.Fprintln(w, "Hello, World!")
	})

	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)

	// key, err := generateKey(32) // AES-256
	// if err != nil {
	// 	log.Fatalf("Failed to generate key: %v", err)
	// }

	// message := "Hello, world!"
	// fmt.Println("Original message:", message)

	// encrypted, err := encrypt(message, key)
	// if err != nil {
	// 	log.Fatalf("Failed to encrypt message: %v", err)
	// }
	// fmt.Println("Encrypted message:", encrypted)
	// // Save the key
	// fmt.Println("Key:", key)

	// decrypted, err := decrypt(encrypted, key)
	// if err != nil {
	// 	log.Fatalf("Failed to decrypt message: %v", err)
	// }
	// fmt.Println("Decrypted message:", decrypted)
}
