package repository

import (
	"github.com/smartik/api/internal/models"
	"gorm.io/gorm"
)

type MemorandumRepository struct {
	db *gorm.DB
}

// Creates a new instance of MemorandumRepository
func NewMemorandumRepository(db *gorm.DB) *MemorandumRepository {
	return &MemorandumRepository{db}
}

// Creates a new memorandum record in the database
func (r *MemorandumRepository) Create(memorandum *models.Memorandum) error {
	return r.db.Create(memorandum).Error
}

// Retrieves all memorandums from the database
func (r *MemorandumRepository) GetAll() (*[]models.Memorandum, error) {
	var memorandums []models.Memorandum
	if err := r.db.Find(&memorandums).Error; err != nil {
		return nil, err
	}
	return &memorandums, nil
}

// Retrieves a specific memorandum by its ID
func (r *MemorandumRepository) GetById(id string) (*models.Memorandum, error) {
	var memorandum models.Memorandum
	if err := r.db.Where("id = ?", id).First(&memorandum).Error; err != nil {
		return nil, err
	}
	return &memorandum, nil
}

// Updates an existing memorandum record
func (r *MemorandumRepository) Update(id string, data *models.Memorandum) (*models.Memorandum, error) {
	memorandum, err := r.GetById(id)
	if err != nil {
		return nil, err
	}

	if err := r.db.Model(&memorandum).Updates(data).Error; err != nil {
		return nil, err
	}
	return memorandum, nil
}

// Deletes a memorandum from the database
func (r *MemorandumRepository) Delete(id string) error {
	memorandum, err := r.GetById(id)
	if err != nil {
		return err
	}
	return r.db.Delete(memorandum).Error
}
