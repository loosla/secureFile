package main

type PasswordRequest struct {
	Password string `json:"password"`
}

type FileResponse struct {
	Password string `json:"password"`
	Content  string `json:"content"`
}
