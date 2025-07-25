package models

import "time"

type Exam struct {
	BaseModel
	Date time.Time `json:"date" gorm:"index:idx_exam_date;not null" validate:"required"`
}

type UpdateExam struct {
	Date *time.Time `json:"date" validate:"omitempty"`
}
