package main

type FilesContentRequest struct {
	Password string `json:"password"`
	File     string `json:"file"`
}

type FilesContentResponse struct {
	Content string `json:"content"`
}

type FilesSaveRequest struct {
	Password string `json:"password"`
	Content  string `json:"content"`
	// TODO: add file
}
