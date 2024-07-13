package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

var (
	mu          sync.RWMutex
	defaultFile = "../files/file.txt" // TODO: receive from user.
)

func filesContentHandler(w http.ResponseWriter, r *http.Request) {
	mu.RLock()
	defer mu.RUnlock()

	var req FilesContentRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println("File: ", req.File)
	fmt.Println("Password: ", req.Password)

	fileContent, err := readFromFile(defaultFile)
	if err != nil {
		http.Error(w, "file not found", http.StatusNotFound)
		return
	}

	key := createAESKey(req.Password)

	decrypted, err := decrypt(fileContent, []byte(key))
	if err != nil {
		log.Fatalf("failed to decrypt message: %v", err)
	}

	resp := FilesContentResponse{Content: string(decrypted)}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func filesSaveHandler(w http.ResponseWriter, r *http.Request) {
	var file FilesSaveRequest
	if err := json.NewDecoder(r.Body).Decode(&file); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	key := createAESKey(file.Password)

	encrypted, err := encrypt(file.Content, []byte(key))
	if err != nil {
		log.Fatalf("failed to encrypt message: %v", err)
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
