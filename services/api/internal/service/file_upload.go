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

func (u *UploadResult) addUploadError(filename, errorMsg string) {
	u.FailedUploads = append(u.FailedUploads, FileUploadError{
		Filename: filename,
		Error:    errorMsg,
	})
}
