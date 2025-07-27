package repository

import (
	"github.com/smartik/api/internal/models"
	"gorm.io/gorm"
)

type StudentRepository struct {
	db *gorm.DB
}

func NewStudentRepository(db *gorm.DB) *StudentRepository {
	return &StudentRepository{db}
}

func (r *StudentRepository) Create(student *models.Student) error {
	return r.db.Create(student).Error
}

func (r *StudentRepository) GetAll() (*[]models.Student, error) {
	var students []models.Student
	if err := r.db.Find(&students).Error; err != nil {
		return nil, err
	}
	return &students, nil
}

func (r *StudentRepository) GetById(id string) (*models.Student, error) {
	var student models.Student
	if err := r.db.Where("id=?", id).First(&student).Error; err != nil {
		return nil, err
	}
	return &student, nil
}

func (r *StudentRepository) Update(id string, data *models.UpdateStudent) (*models.Student, error) {
	student, err := r.GetById(id)
	if err != nil {
		return nil, err
	}

	if err := r.db.Model(&student).Updates(data).Error; err != nil {
		return nil, err
	}

	return student, nil
}

func (r *StudentRepository) Delete(id string) error {
	student, err := r.GetById(id)
	if err != nil {
		return err
	}

	return r.db.Delete(&student).Error
}
