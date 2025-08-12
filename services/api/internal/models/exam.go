package models

import "time"

type Exam struct {
	BaseModel
	Date          time.Time      `json:"date" gorm:"index:idx_exam_date;not null" validate:"required"`
	TotalMarks    int            `json:"total_marks" gorm:"type:int;default:1;not null" validate:"numeric,min=0"`
	AnswerScripts []AnswerScript `json:"answer_scripts,omitempty" gorm:"foreignKey:ExamId;references:Id;constraint:OnDelete:CASCADE" validate:"-"`
}

type UpdateExam struct {
	Date          *time.Time `json:"date" validate:"omitempty"`
	TotalMarks    *int       `json:"total_marks,omitempty" validate:"omitempty,numeric,min=0"`
	AnswerScripts *[]string  `json:"answer_scripts,omitempty" validate:"omitempty"` // IDs of answer scripts to add to attach to the exam
}
