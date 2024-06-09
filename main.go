package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
)

var (
	mu          sync.Mutex
	defaultFile = "files/file.txt"
)

type PasswordRequest struct {
	Password string `json:"password"`
}

type FileResponse struct {
	Password string `json:"password"`
	Content  string `json:"content"`
}

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

	_, err = file.Write([]byte(data))
	if err != nil {
		return fmt.Errorf("error writing to file: %v", err)
	}

	return os.WriteFile(fileName, []byte(data), 0644)
}

func getTextHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	var req PasswordRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	fileContent, err := readFromFile(defaultFile)
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	key := createAESKey(req.Password)

	decrypted, err := decrypt(fileContent, []byte(key))
	if err != nil {
		log.Fatalf("Failed to decrypt message: %v", err)
	}

	resp := FileResponse{Content: string(decrypted)}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func saveFileHandler(w http.ResponseWriter, r *http.Request) {
	var response FileResponse
	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	key := createAESKey(response.Password)

	encrypted, err := encrypt(response.Content, []byte(key))
	if err != nil {
		log.Fatalf("Failed to encrypt message: %v", err)
	}

	mu.Lock()
	err = writeToFile(defaultFile, encrypted)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	mu.Unlock()

	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/api/get-file", getTextHandler)
	http.HandleFunc("/api/save-file", saveFileHandler)

	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
