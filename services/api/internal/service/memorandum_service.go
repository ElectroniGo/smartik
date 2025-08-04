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

type MemorandumService struct {
	repo        *repository.MemorandumRepository
	minioClient *minio.Client
	cfg         *config.Env
}

type MemorandumUploadResult struct {
	UploadResult
	SuccessfulUploads []models.Memorandum `json:"successful_uploads"`
}

// Creates a new instance of MemorandumService
func NewMemorandumService(
	repo *repository.MemorandumRepository,
	minioClient *minio.Client,
	cfg *config.Env,
) *MemorandumService {
	return &MemorandumService{repo, minioClient, cfg}
}

// Handles the upload of a single memorandum file
func (s *MemorandumService) UploadFile(file *multipart.FileHeader, examId string, result *MemorandumUploadResult) (*MemorandumUploadResult, error) {
	src, err := file.Open()
	if err != nil {
		result.FailedUploads = append(result.FailedUploads, FileUploadError{
			Filename: file.Filename,
			Error:    "Failed to open file: " + err.Error(),
		})
		return nil, err
	}
	defer src.Close()

	// Upload the file to MinIO
	if err := s.uploadToStorage(file.Filename, src, file.Size, file.Header.Get("Content-Type")); err != nil {
		result.FailedUploads = append(result.FailedUploads, FileUploadError{
			Filename: file.Filename,
			Error:    "Failed to upload file to storage: " + err.Error(),
		})
	}

	// Create database record
	memorandum := &models.Memorandum{
		FileName: file.Filename,
		ExamId:   examId,
	}

	if err := s.repo.Create(memorandum); err != nil {
		if deleteErr := s.repo.Delete(file.Filename); deleteErr != nil {
			log.Errorf("Failed to delete memorandum record after upload error: %v", deleteErr)
		}

		result.FailedUploads = append(result.FailedUploads, FileUploadError{
			Filename: file.Filename,
		})
	}

	result.SuccessfulUploads = append(result.SuccessfulUploads, *memorandum)
	return result, nil
}

// Retrieves a file stream for serving the memorandum file
func (s *MemorandumService) GetFileStream(id string) (*FileStreamResult, error) {
	memorandum, err := s.repo.GetById(id)
	if err != nil {
		return nil, err
	}

	// Get the file from MinIO storage
	object, err := s.minioClient.GetObject(
		context.Background(),
		s.cfg.MinioStorageBucket,
		memorandum.FileName,
		minio.GetObjectOptions{},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get file from storage: %v", err)
	}

	// Get file metadata
	objectInfo, err := object.Stat()
	if err != nil {
		object.Close()
		return nil, fmt.Errorf("failed to get file info: %v", err)
	}

	return &FileStreamResult{
		Content:     object,
		ContentType: objectInfo.ContentType,
		Filename:    memorandum.FileName,
		Size:        objectInfo.Size,
	}, nil
}

// Retrieves all memorandums from the database
func (s *MemorandumService) GetAll() (*[]models.Memorandum, error) {
	return s.repo.GetAll()
}

// Retrieves a specific memorandum by its ID
func (s *MemorandumService) GetById(id string) (*models.Memorandum, error) {
	return s.repo.GetById(id)
}

// Removes a memorandum from the database and MinIO storage
func (s *MemorandumService) Delete(id string) error {
	memorandum, err := s.repo.GetById(id)
	if err != nil {
		return err
	}

	// Delete from MinIO storage
	if err := s.deleteFromStorage(memorandum.FileName); err != nil {
		return err
	}

	// Delete from database
	return s.repo.Delete(id)
}

// Handles the actual file upload to MinIO
func (s *MemorandumService) uploadToStorage(filename string, src io.Reader, size int64, contentType string) error {
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
func (s *MemorandumService) deleteFromStorage(filename string) error {
	return s.minioClient.RemoveObject(
		context.Background(),
		s.cfg.MinioStorageBucket,
		filename,
		minio.RemoveObjectOptions{},
	)
}
