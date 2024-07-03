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

func TestFilesContentHandler(t *testing.T) {
	setup()
	defer teardown()

	// Set up the initial content for the file
	password := "testpassword"
	content := "This is a test file content."
	key := createAESKey(password)
	encryptedContent, _ := encrypt(content, key)
	os.WriteFile(testFile, []byte(encryptedContent), 0644)

	// Create a request to pass to the handler
	reqBody := FilesContentRequest{Password: password} // TODO: Should be protected.
	reqBodyBytes, _ := json.Marshal(reqBody)
	req, err := http.NewRequest("POST", "/files/content", bytes.NewBuffer(reqBodyBytes))
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(filesContentHandler)

	// Call the handler
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// TODO: uncomment when can set path to file.
	// Check the response body is what we expect
	// expected := FilesContentResponse{Content: content}
	// var resp FilesContentResponse
	// if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
	// 	t.Fatalf("Could not decode response: %v", err)
	// }
	// if resp.Content != expected.Content {
	// 	t.Errorf("Handler returned unexpected body: got %v want %v",
	// 		resp.Content, expected.Content)
	// }
}

func TestFilesSaveHandler(t *testing.T) {
	setup()
	defer teardown()

	// Create a request to pass to the handler
	password := "testpassword"
	content := "This is a test file content."
	reqBody := FilesSaveRequest{Password: password, Content: content}
	reqBodyBytes, _ := json.Marshal(reqBody)
	req, err := http.NewRequest("POST", "/files/save", bytes.NewBuffer(reqBodyBytes))
	if err != nil {
		t.Fatalf("Could not save request: %v", err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(filesSaveHandler)

	// Call the handler
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// TODO: uncomment when can set path to file.
	// Check if the content was correctly saved to the file
	// fileContent, err := readFromFile("../files/file.txt") // TODO: fix!
	// if err != nil {
	// 	t.Fatalf("Could not read from file: %v", err)
	// }

	// key := createAESKey(password)
	// decryptedContent, err := decrypt(fileContent, key)
	// if err != nil {
	// 	t.Fatalf("Could not decrypt file content: %v", err)
	// }

	// if decryptedContent != content {
	// 	t.Errorf("File content mismatch: got %v want %v", decryptedContent, content)
	// }
}
