package models

type Student struct {
	BaseModel
	FirstName  string `json:"first_name" gorm:"type:varchar(50)" validate:"omitempty,min=3,max=50"`
	LastName   string `json:"last_name" gorm:"type:varchar(50)" validate:"omitempty,min=3,max=50"`
	ExamNumber string `json:"exam_number" gorm:"type:varchar(20);uniqueIndex;not null" validate:"required,min=4,max=20"`
}

type UpdateStudent struct {
	FirstName  *string `json:"first_name,omitempty" validate:"omitempty,min=3,max=50"`
	LastName   *string `json:"last_name,omitempty" validate:"omitempty,min=3,max=50"`
	ExamNumber *string `json:"exam_number,omitempty" validate:"omitempty,min=4,max=20"`
}
