package service

import (
	"github.com/smartik/api/internal/models"
	"github.com/smartik/api/internal/repository"
)

// Handles business logic for student operations
type StudentService struct {
	repo *repository.StudentRepository
}

// Creates a new instance of StudentService
func NewStudentService(repo *repository.StudentRepository) *StudentService {
	return &StudentService{
		repo: repo,
	}
}

// Creates a new student record in the database
func (s *StudentService) Create(student *models.Student) error {
	return s.repo.Create(student)
}

// Retrieves all students from the database
func (s *StudentService) GetAll() (*[]models.Student, error) {
	return s.repo.GetAll()
}

// Retrieves a specific student by their ID
func (s *StudentService) GetById(id string) (*models.Student, error) {
	return s.repo.GetById(id)
}

// Modifies an existing student record
func (s *StudentService) Update(id string, updateData *models.UpdateStudent) (*models.Student, error) {
	return s.repo.Update(id, updateData)
}

// Removes a student from the database
func (s *StudentService) Delete(id string) error {
	return s.repo.Delete(id)
}
