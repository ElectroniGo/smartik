package repository

import (
	"github.com/smartik/api/internal/models"
	"gorm.io/gorm"
)

type AnswerScriptRepository struct {
	db *gorm.DB
}

// Creates a new instance of AnswerScriptRepository
func NewAnswerScriptRepository(db *gorm.DB) *AnswerScriptRepository {
	return &AnswerScriptRepository{db}
}

// Creates a new answer script record in the database
func (r *AnswerScriptRepository) Create(answerScript *models.AnswerScript) error {
	return r.db.Create(answerScript).Error
}

// Retrieves all answer scripts from the database
func (r *AnswerScriptRepository) GetAll() (*[]models.AnswerScript, error) {
	var answerScripts []models.AnswerScript
	if err := r.db.Find(&answerScripts).Error; err != nil {
		return nil, err
	}
	return &answerScripts, nil
}

// Retrieves a specific answer script by its ID
func (r *AnswerScriptRepository) GetById(id string) (*models.AnswerScript, error) {
	var answerScript models.AnswerScript
	if err := r.db.Where("id = ?", id).First(&answerScript).Error; err != nil {
		return nil, err
	}
	return &answerScript, nil
}

// Updates an existing answer script record
func (r *AnswerScriptRepository) Update(id string, data *models.AnswerScript) (*models.AnswerScript, error) {
	answerScript, err := r.GetById(id)
	if err != nil {
		return nil, err
	}

	if err := r.db.Model(&answerScript).Updates(data).Error; err != nil {
		return nil, err
	}
	return answerScript, nil
}

// Deletes an answer script from the database
func (r *AnswerScriptRepository) Delete(id string) error {
	answerScript, err := r.GetById(id)
	if err != nil {
		return err
	}
	return r.db.Delete(answerScript).Error
}
