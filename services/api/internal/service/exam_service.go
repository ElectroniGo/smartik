package service

import (
	"github.com/smartik/api/internal/models"
	"github.com/smartik/api/internal/repository"
)

// Handles business logic for exam operations
type ExamService struct {
	repo *repository.ExamRepository
}

// Creates a new instance of ExamService
func NewExamService(repo *repository.ExamRepository) *ExamService {
	return &ExamService{
		repo: repo,
	}
}

// Creates a new exam record in the database
func (s *ExamService) Create(exam *models.Exam) error {
	return s.repo.Create(exam)
}

// Retrieves all exams from the database
func (s *ExamService) GetAll() (*[]models.Exam, error) {
	return s.repo.GetAll()
}

// Retrieves a specific exam by its ID
func (s *ExamService) GetById(id string) (*models.Exam, error) {
	return s.repo.GetById(id)
}

// Modifies an existing exam record
func (s *ExamService) Update(id string, updateData *models.UpdateExam) (*models.Exam, error) {
	return s.repo.Update(id, updateData)
}

// Removes an exam from the database
func (s *ExamService) Delete(id string) error {
	return s.repo.Delete(id)
}
