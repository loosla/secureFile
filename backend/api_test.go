package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"sync"
	"testing"
)

const (
	testFile = "testfile.txt"
)

var mutex sync.Mutex

func setup() {
	mutex.Lock()
	os.WriteFile(testFile, []byte("setup"), 0644)
	mutex.Unlock()
}

func teardown() {
	mutex.Lock()
	os.Remove(testFile)
	mutex.Unlock()
}

func TestGetFileHandler(t *testing.T) {
	setup()
	defer teardown()

	// Set up the initial content for the file
	password := "testpassword"
	content := "This is a test file content."
	key := createAESKey(password)
	encryptedContent, _ := encrypt(content, key)
	os.WriteFile(testFile, []byte(encryptedContent), 0644)

	// Create a request to pass to the handler
	reqBody := Password{Password: password}
	reqBodyBytes, _ := json.Marshal(reqBody)
	req, err := http.NewRequest("POST", "/api/get-file", bytes.NewBuffer(reqBodyBytes))
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getFileHandler)

	// Call the handler
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect
	expected := File{Content: content}
	var resp File
	if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
		t.Fatalf("Could not decode response: %v", err)
	}
	if resp.Content != expected.Content {
		t.Errorf("Handler returned unexpected body: got %v want %v",
			resp.Content, expected.Content)
	}
}

func TestSaveFileHandler(t *testing.T) {
	setup()
	defer teardown()

	// Create a request to pass to the handler
	password := "testpassword"
	content := "This is a test file content."
	reqBody := File{Password: password, Content: content}
	reqBodyBytes, _ := json.Marshal(reqBody)
	req, err := http.NewRequest("POST", "/api/save-file", bytes.NewBuffer(reqBodyBytes))
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(saveFileHandler)

	// Call the handler
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check if the content was correctly saved to the file
	fileContent, err := readFromFile("../files/file.txt") // TODO: fix!
	if err != nil {
		t.Fatalf("Could not read from file: %v", err)
	}

	key := createAESKey(password)
	decryptedContent, err := decrypt(fileContent, key)
	if err != nil {
		t.Fatalf("Could not decrypt file content: %v", err)
	}

	if decryptedContent != content {
		t.Errorf("File content mismatch: got %v want %v", decryptedContent, content)
	}
}
