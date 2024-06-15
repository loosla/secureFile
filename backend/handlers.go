package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func getFileHandler(w http.ResponseWriter, r *http.Request) {
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
