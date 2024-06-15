package main

type Password struct {
	Password string `json:"password"`
}

type File struct {
	Password string `json:"password"`
	Content  string `json:"content"`
}
