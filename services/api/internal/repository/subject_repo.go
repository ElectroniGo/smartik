package repository

import (
	"github.com/smartik/api/internal/models"
	"gorm.io/gorm"
)

type SubjectRepository struct {
	db *gorm.DB
}

func NewSubjectRepository(db *gorm.DB) *SubjectRepository {
	return &SubjectRepository{db}
}

func (r *SubjectRepository) Create(subject *models.Subject) error {
	return r.db.Create(subject).Error
}

func (r *SubjectRepository) GetAll() (*[]models.Subject, error) {
	var subjects []models.Subject
	if err := r.db.Find(&subjects).Error; err != nil {
		return nil, err
	}

	return &subjects, nil
}

func (r *SubjectRepository) GetById(id string) (*models.Subject, error) {
	var subject models.Subject
	if err := r.db.Where("id=?", id).First(&subject).Error; err != nil {
		return nil, err
	}

	return &subject, nil
}

func (r *SubjectRepository) Update(id string, data *models.UpdateSubject) (*models.Subject, error) {
	var subject models.Subject
	if err := r.db.Where("id = ?", id).First(&subject).Error; err != nil {
		return nil, err
	}

	if err := r.db.Model(&subject).Updates(data).Error; err != nil {
		return nil, err
	}

	return &subject, nil
}

func (r *SubjectRepository) Delete(id string) error {
	var subject models.Subject
	if err := r.db.Where("id = ?", id).First(&subject).Error; err != nil {
		return err
	}

	return r.db.Delete(&subject).Error
}
