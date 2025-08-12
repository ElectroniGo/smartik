package repository

import (
	"github.com/smartik/api/internal/models"
	"gorm.io/gorm"
)

type SubjectRepository struct {
	db *gorm.DB
}

// Creates a new instance of SubjectRepository
func NewSubjectRepository(db *gorm.DB) *SubjectRepository {
	return &SubjectRepository{db}
}

// Creates a new subject record in the database
func (r *SubjectRepository) Create(subject *models.Subject) error {
	return r.db.Create(subject).Error
}

// Retrieves all subjects from the database
func (r *SubjectRepository) GetAll() (*[]models.Subject, error) {
	var subjects []models.Subject
	if err := r.db.Find(&subjects).Error; err != nil {
		return nil, err
	}

	return &subjects, nil
}

// Retrieves a specific subject by its ID
func (r *SubjectRepository) GetById(id string) (*models.Subject, error) {
	var subject models.Subject
	if err := r.db.Where("id=?", id).First(&subject).Error; err != nil {
		return nil, err
	}

	return &subject, nil
}

// Updates an existing subject record
func (r *SubjectRepository) Update(id string, data *models.UpdateSubject) (*models.Subject, error) {
	subject, err := r.GetById(id)
	if err != nil {
		return nil, err
	}

	if err := r.db.Model(&subject).Updates(data).Error; err != nil {
		return nil, err
	}

	return subject, nil
}

// Deletes a subject from the database
func (r *SubjectRepository) Delete(id string) error {
	subject, err := r.GetById(id)
	if err != nil {
		return err
	}

	return r.db.Delete(&subject).Error
}
