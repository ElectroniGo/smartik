package repository

import (
	"github.com/smartik/api/internal/models"
	"gorm.io/gorm"
)

type AnswerScriptRepository struct {
	db *gorm.DB
}

func NewAnswerScriptRepository(db *gorm.DB) *AnswerScriptRepository {
	return &AnswerScriptRepository{db}
}

func (r *AnswerScriptRepository) Create(answerScript *models.AnswerScript) error {
	if err := r.db.Create(answerScript).Error; err != nil {
		return err
	}

	return nil
}

func (r *AnswerScriptRepository) GetAll() (*[]models.AnswerScript, error) {
	var answerScripts []models.AnswerScript
	if err := r.db.Find(&answerScripts).Error; err != nil {
		return nil, err
	}
	return &answerScripts, nil
}

func (r *AnswerScriptRepository) GetById(id string) (*models.AnswerScript, error) {
	var answerScript models.AnswerScript
	if err := r.db.Where("id = ?", id).First(&answerScript).Error; err != nil {
		return nil, err
	}
	return &answerScript, nil
}

func (r *AnswerScriptRepository) Update(id string, data *models.AnswerScript) (*models.AnswerScript, error) {
	answerScript, err := r.GetById(id)
	if err != nil {
		return nil, err
	}

	if err := r.db.Model(&answerScript).Updates(&data).Error; err != nil {
		return nil, err
	}
	return answerScript, nil
}

func (r *AnswerScriptRepository) Delete(id string) error {
	answerScript, err := r.GetById(id)
	if err != nil {
		return err
	}
	return r.db.Delete(answerScript).Error
}
