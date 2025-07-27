package models

import "time"

type ProcessingStatus string

const (
	StatusProcessing ProcessingStatus = "processing"
	StatusUploaded   ProcessingStatus = "uploaded"
	StatusFailed     ProcessingStatus = "failed"
)

type AnswerScript struct {
	BaseModel
	FileName           string           `json:"file_name" gorm:"type:varchar(255);not null" validate:"required,min=3,max=255"`
	FileUrl            *string          `json:"file_url" gorm:"type:text" validate:"omitempty,url"`
	StudentId          *string          `json:"student_id" gorm:"type:varchar(25)" validate:"omitempty"`
	Student            *Student         `json:"student,omitempty" gorm:"foreignKey:StudentId;references:Id;constraint:OnDelete:SET NULL" validate:"-"`
	SubjectId          *string          `json:"subject_id" gorm:"type:varchar(25)" validate:"omitempty"`
	Subject            *Subject         `json:"subject,omitempty" gorm:"foreignKey:SubjectId;references:Id;constraint:OnDelete:SET NULL" validate:"-"`
	ExamId             *string          `json:"exam_id" gorm:"type:varchar(25)" validate:"omitempty"`
	Exam               *Exam            `json:"exam,omitempty" gorm:"foreignKey:ExamId;references:Id;constraint:OnDelete:CASCADE" validate:"-"`
	TotalMarks         *int             `json:"total_marks" gorm:"type:int;default:NULL" validate:"omitempty,numeric,min=0"`
	MaxMarks           *int             `json:"max_marks" gorm:"type:int;default:NULL" validate:"omitempty,numeric,min=0"`
	ScannedExamNumber  *string          `json:"scanned_exam_number" gorm:"type:varchar(20)" validate:"omitempty,min=4,max=20"`
	Status             ProcessingStatus `json:"processing_status" gorm:"type:varchar(20);default:processing" validate:"omitempty,oneof=processing uploaded failed"` // can be 'processing', 'uploaded', or 'failed'
	MatchedAt          *time.Time       `json:"matched_at" gorm:"type:timestamp;default:NULL" validate:"omitempty"`
	MatchingConfidence *float32         `json:"matching_confidence" gorm:"type:float" validate:"omitempty,numeric"` // Confidence interval for the OCR extracted scanned exam number
}

type UpdateAnswerScript struct {
	FileName           *string           `json:"file_name,omitempty" validate:"omitempty,min=3,max=255"`
	FileUrl            *string           `json:"file_url,omitempty" validate:"omitempty,url"`
	StudentId          *string           `json:"student_id,omitempty" validate:"omitempty"`
	SubjectId          *string           `json:"subject_id,omitempty" validate:"omitempty"`
	ExamId             *string           `json:"exam_id,omitempty" validate:"omitempty"`
	TotalMarks         *int              `json:"total_marks,omitempty" validate:"omitempty,numeric,min=0"`
	MaxMarks           *int              `json:"max_marks,omitempty" validate:"omitempty,numeric,min=0"`
	ScannedExamNumber  *string           `json:"scanned_exam_number,omitempty" validate:"omitempty,min=4,max=20"`
	Status             *ProcessingStatus `json:"processing_status,omitempty" validate:"omitempty,oneof=processing uploaded failed"` // can be 'processing', 'uploaded', or 'failed'
	MatchingConfidence *float32          `json:"matching_confidence,omitempty" validate:"omitempty,numberic"`
	MatchedAt          *time.Time        `json:"matched_at,omitempty" validate:"omitempty"`
}
