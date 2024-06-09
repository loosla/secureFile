package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
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

type Response struct {
	Text string `json:"text"`
}

var (
	mu          sync.Mutex
	defaultFile = "files/file.txt"
)

func readFromFile(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func writeToFile(fileName, data string) error {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	// Write the data to the file
	_, err = file.Write([]byte(data))
	if err != nil {
		return fmt.Errorf("error writing to file: %v", err)
	}

	return os.WriteFile(fileName, []byte(data), 0644)
}

func getTextHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	textData, err := readFromFile(defaultFile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := Response{Text: textData}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func saveTextHandler(w http.ResponseWriter, r *http.Request) {
	var response Response
	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mu.Lock()
	err := writeToFile(defaultFile, response.Text)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	mu.Unlock()

	w.WriteHeader(http.StatusOK)
}

func main() {

	http.HandleFunc("/api/text", getTextHandler)
	http.HandleFunc("/api/save", saveTextHandler)

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
