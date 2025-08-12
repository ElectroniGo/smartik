package service

import (
	"github.com/smartik/api/internal/models"
	"github.com/smartik/api/internal/repository"
)

// Handles business logic for subject operations
type SubjectService struct {
	repo *repository.SubjectRepository
}

// Creates a new instance of SubjectService
func NewSubjectService(repo *repository.SubjectRepository) *SubjectService {
	return &SubjectService{
		repo: repo,
	}
}

// Creates a new subject record in the database
func (s *SubjectService) Create(subject *models.Subject) error {
	return s.repo.Create(subject)
}

// Retrieves all subjects from the database
func (s *SubjectService) GetAll() (*[]models.Subject, error) {
	return s.repo.GetAll()
}

// Retrieves a specific subject by its ID
func (s *SubjectService) GetById(id string) (*models.Subject, error) {
	return s.repo.GetById(id)
}

// Modifies an existing subject record
func (s *SubjectService) Update(id string, updateData *models.UpdateSubject) (*models.Subject, error) {
	return s.repo.Update(id, updateData)
}

// Removes a subject from the database
func (s *SubjectService) Delete(id string) error {
	return s.repo.Delete(id)
}
