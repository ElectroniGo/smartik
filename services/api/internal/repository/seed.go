package repository

import (
	"time"

	"github.com/smartik/api/internal/models"
	"gorm.io/gorm"
)

func SeedDatabase(db *gorm.DB) error {
	var count int64
	if err := db.Model(&models.Student{}).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return nil // Database already seeded
	}

	// Seed initial data
	students := []models.Student{
		{FirstName: "John", LastName: "Doe", ExamNumber: "JOH5196"},
		{FirstName: "Jane", LastName: "Dwayne", ExamNumber: "JAN5196"},
		{FirstName: "Alice", LastName: "Smith", ExamNumber: "ALI5196"},
	}

	for _, student := range students {
		if err := db.Create(&student).Error; err != nil {
			return err
		}
	}

	subjects := []models.Subject{
		{Name: "Mathematics", Code: "MATH101"},
		{Name: "Physics", Code: "PHY101"},
		{Name: "Chemistry", Code: "CHEM101"},
	}

	for _, subject := range subjects {
		if err := db.Create(&subject).Error; err != nil {
			return err
		}
	}

	exams := []models.Exam{
		{Date: time.Now().Add(24 * 30 * time.Hour)},  // 30 days from now
		{Date: time.Now().Add(24 * 60 * time.Hour)},  // 60 days from now
		{Date: time.Now().Add(24 * 120 * time.Hour)}, // 120 days from now
	}

	for _, exam := range exams {
		if err := db.Create(&exam).Error; err != nil {
			return err
		}
	}

	return nil
}
