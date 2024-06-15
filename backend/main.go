package main

import (
	"fmt"
	"net/http"
	"sync"
)

var (
	mu          sync.Mutex
	defaultFile = "../files/file.txt"
)

func main() {
	http.HandleFunc("/api/get-file", getFileHandler)
	http.HandleFunc("/api/save-file", saveFileHandler)

	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
