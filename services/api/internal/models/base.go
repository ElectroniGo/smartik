package models

import (
	"time"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type BaseModel struct {
	Id        string    `json:"id" gorm:"primaryKey;type:varchar(25)"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (b *BaseModel) BeforeCreate(tx *gorm.DB) error {
	if b.Id == "" {
		err := SetId(&b.Id)
		if err != nil {
			return err
		}
	}
	return nil
}

func (b *BaseModel) BeforeUpdate(tx *gorm.DB) error {
	tx.Statement.Omit("id", "created_at", "updated_at")
	return nil
}

func SetId(id *string) error {
	newId, err := gonanoid.New(21)
	if err != nil {
		return err
	}
	*id = newId
	return nil
}
