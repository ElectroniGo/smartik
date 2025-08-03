package service

import "io"

type UploadResult struct {
	FailedUploads []FileUploadError `json:"failed_uploads"`
}

type FileUploadError struct {
	Filename string
	Error    string
}

type FileStreamResult struct {
	Content     io.ReadCloser
	ContentType string
	Filename    string
	Size        int64
}
