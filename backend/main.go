package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/file", getFileHandler)
	http.HandleFunc("/file/update", updateFileHandler)

	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
