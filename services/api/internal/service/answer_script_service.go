package service

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"

	"github.com/labstack/gommon/log"
	minio "github.com/minio/minio-go/v7"
	"github.com/smartik/api/internal/config"
	"github.com/smartik/api/internal/models"
	"github.com/smartik/api/internal/repository"
)

// Handles business logic for answer script operations
type AnswerScriptService struct {
	repo        *repository.AnswerScriptRepository
	minioClient *minio.Client
	cfg         *config.Env
}

type AnswerScriptUploadResult struct {
	UploadResult
	SuccessfulUploads []models.AnswerScript `json:"successful_uploads"`
}

// Creates a new instance of AnswerScriptService
func NewAnswerScriptService(
	repo *repository.AnswerScriptRepository,
	minioClient *minio.Client,
	cfg *config.Env,
) *AnswerScriptService {
	return &AnswerScriptService{
		repo:        repo,
		minioClient: minioClient,
		cfg:         cfg,
	}
}

// Handles the upload of multiple answer script files
// Processes each file individually and returns a summary of successes and failures
func (s *AnswerScriptService) UploadFiles(files []*multipart.FileHeader) (*AnswerScriptUploadResult, error) {
	result := &AnswerScriptUploadResult{
		SuccessfulUploads: []models.AnswerScript{},
		UploadResult: UploadResult{
			FailedUploads: []FileUploadError{},
		},
	}

	// Process each file individually
	for _, file := range files {
		if err := s.uploadSingleFile(file, result); err != nil {
			continue // error handled by in `uploadSingleFile
		}
	}

	return result, nil
}

// Processes a single file upload with proper error handling and rollback
func (s *AnswerScriptService) uploadSingleFile(file *multipart.FileHeader, result *AnswerScriptUploadResult) error {
	src, err := file.Open()
	if err != nil {
		s.addUploadError(result, file.Filename, "Failed to open file: "+err.Error())
		return err
	}
	defer src.Close()

	// Upload file to MinIO storage
	if err := s.uploadToStorage(file.Filename, src, file.Size, file.Header.Get("Content-Type")); err != nil {
		s.addUploadError(result, file.Filename, "Failed to upload to storage: "+err.Error())
		return err
	}

	// Create database record for the uploaded file
	answerScript := &models.AnswerScript{
		FileName: file.Filename,
		Status:   models.StatusUploaded,
	}

	if err := s.repo.Create(answerScript); err != nil {
		// Deletes file from MinIO if database save fails
		if deleteErr := s.deleteFromStorage(file.Filename); deleteErr != nil {
			log.Errorf("Failed to rollback file deletion for %s: %v", file.Filename, deleteErr)
		}

		s.addUploadError(result, file.Filename, "Failed to save to database: "+err.Error())
		return err
	}

	// Add successful upload to result
	result.SuccessfulUploads = append(result.SuccessfulUploads, *answerScript)
	return nil
}

// Handles the actual file upload to MinIO
func (s *AnswerScriptService) uploadToStorage(filename string, src io.Reader, size int64, contentType string) error {
	_, err := s.minioClient.PutObject(
		context.Background(),
		s.cfg.MinioStorageBucket,
		filename,
		src,
		size,
		minio.PutObjectOptions{ContentType: contentType},
	)
	return err
}

// Removes a file from MinIO storage
func (s *AnswerScriptService) deleteFromStorage(filename string) error {
	return s.minioClient.RemoveObject(
		context.Background(),
		s.cfg.MinioStorageBucket,
		filename,
		minio.RemoveObjectOptions{},
	)
}

// Helper method to add upload errors to the result
func (s *AnswerScriptService) addUploadError(result *AnswerScriptUploadResult, filename, errorMsg string) {
	result.FailedUploads = append(result.FailedUploads, FileUploadError{
		Filename: filename,
		Error:    errorMsg,
	})
}

// Retrieves all answer scripts from the database
func (s *AnswerScriptService) GetAll() (*[]models.AnswerScript, error) {
	return s.repo.GetAll()
}

// Retrieves a specific answer script by its ID
func (s *AnswerScriptService) GetById(id string) (*models.AnswerScript, error) {
	return s.repo.GetById(id)
}

// Modifies an existing answer script record
func (s *AnswerScriptService) Update(id string, data *models.AnswerScript) (*models.AnswerScript, error) {
	return s.repo.Update(id, data)
}

// Removes an answer script from both database and storage
func (s *AnswerScriptService) Delete(id string) error {
	answerScript, err := s.repo.GetById(id)
	if err != nil {
		return err
	}

	// Delete from storage (Minio)
	if err := s.deleteFromStorage(answerScript.FileName); err != nil {
		return err
	}

	// Delete from database
	if err := s.repo.Delete(id); err != nil {
		return err
	}

	return nil
}

// Retrieves a file stream from storage for serving files
func (s *AnswerScriptService) GetFileStream(id string) (*FileStreamResult, error) {
	answerScript, err := s.repo.GetById(id)
	if err != nil {
		return nil, err
	}

	// Get file object from MinIO
	object, err := s.minioClient.GetObject(
		context.Background(),
		s.cfg.MinioStorageBucket,
		answerScript.FileName,
		minio.GetObjectOptions{},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get file from storage: %w", err)
	}

	// Get file metadata
	objectInfo, err := object.Stat()
	if err != nil {
		object.Close()
		return nil, fmt.Errorf("failed to get file info: %w", err)
	}

	return &FileStreamResult{
		Content:     object,
		ContentType: objectInfo.ContentType,
		Filename:    answerScript.FileName,
		Size:        objectInfo.Size,
	}, nil
}
