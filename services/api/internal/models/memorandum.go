package models

type Memorandum struct {
	BaseModel
	FileName string `json:"file_name" gorm:"type:varchar(255);not null" validate:"omitempty"`
	ExamId   string `json:"exam_id" gorm:"type:varchar(25);not null" validate:"required"`
	Exam     *Exam  `json:"exam,omitempty" gorm:"foreignKey:ExamId;references:Id;constraint:OnDelete:CASCADE"`
}
