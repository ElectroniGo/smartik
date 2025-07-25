package models

type Subject struct {
	BaseModel
	Name        string `json:"name" gorm:"type:varchar(100);not null" validate:"required,min=3,max=100"`
	Code        string `json:"code" gorm:"type:varchar(10);not null" validate:"required,min=2,max=10"`
	Description string `json:"description" gorm:"type:text" validate:"omitempty,max=500"`
}

type UpdateSubject struct {
	Name        *string `json:"name" validate:"omitempty,min=3,max=100"`
	Code        *string `json:"code" validate:"omitempty,min=2,max=10"`
	Description *string `json:"description" validate:"omitempty,max=500"`
}
