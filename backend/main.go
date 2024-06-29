package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/files/content", filesContentHandler)
	http.HandleFunc("/files/save", filesSaveHandler)

	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
