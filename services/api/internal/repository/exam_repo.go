package repository

import (
	"github.com/smartik/api/internal/models"
	"gorm.io/gorm"
)

type ExamRepository struct {
	db *gorm.DB
}

func NewExamRepository(db *gorm.DB) *ExamRepository {
	return &ExamRepository{db}
}

func (r *ExamRepository) Create(exam *models.Exam) error {
	return r.db.Create(exam).Error
}

func (r *ExamRepository) GetAll() (*[]models.Exam, error) {
	var exams []models.Exam
	if err := r.db.Find(&exams).Error; err != nil {
		return nil, err
	}
	return &exams, nil
}

func (r *ExamRepository) GetById(id string) (*models.Exam, error) {
	var exam models.Exam
	if err := r.db.Where("id = ?", id).First(&exam).Error; err != nil {
		return nil, err
	}

	return &exam, nil
}

func (r *ExamRepository) Update(id string, data *models.UpdateExam) (*models.Exam, error) {
	var exam models.Exam
	if err := r.db.Where("id = ?", id).First(&exam).Error; err != nil {
		return nil, err
	}

	if err := r.db.Model(&exam).Updates(&data).Error; err != nil {
		return nil, err
	}

	return &exam, nil
}

func (r *ExamRepository) Delete(id string) error {
	var exam models.Exam
	if err := r.db.Where("id = ?", id).First(&exam).Error; err != nil {
		return err
	}
	return r.db.Delete(&exam).Error
}
