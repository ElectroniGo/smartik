package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/smartik/api/internal/models"
	"github.com/smartik/api/internal/repository"
	"gorm.io/gorm"
)

type StudentHandler struct {
	studentRepo *repository.StudentRepository
}

func NewStudentHandler(repo *repository.StudentRepository) *StudentHandler {
	return &StudentHandler{repo}
}

func (h *StudentHandler) CreateStudent(c echo.Context) error {
	var newStudent models.Student
	if err := c.Bind(&newStudent); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid input",
			"error":   err.Error(),
		})
	}

	if err := c.Validate(&newStudent); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Validation failed",
			"errors":  err.Error(),
		})
	}

	if err := h.studentRepo.Create(&newStudent); err != nil {
		log.Errorf("Failed to create student: %v", err)

		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to create student",
		})
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "Student created successfully",
		"student": newStudent,
	})
}

func (h *StudentHandler) GetAllStudents(c echo.Context) error {
	students, err := h.studentRepo.GetAll()
	if err != nil {
		log.Errorf("Failed to retrieve students: %v", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to retrieve students",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message":  "Students retrieved successfully",
		"students": students,
	})
}

func (h *StudentHandler) GetStudentById(c echo.Context) error {
	id := c.Param("id")

	student, err := h.studentRepo.GetById(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, echo.Map{
				"message": "Student not found",
			})
		}

		log.Errorf("Failed to get student by exam number: %v", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to retrieve student",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Student retrieved successfully",
		"student": student,
	})
}

func (h *StudentHandler) UpdateStudent(c echo.Context) error {
	id := c.Param("id")
	var updateData models.UpdateStudent

	if err := c.Bind(&updateData); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid input",
			"error":   err.Error(),
		})
	}

	if err := c.Validate(&updateData); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Validation failed",
			"errors":  err.Error(),
		})
	}

	updatedStudent, err := h.studentRepo.Update(id, &updateData)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, echo.Map{
				"message": "Student not found",
			})
		}

		log.Errorf("Failed to update student: %v", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to update student",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Student updated successfully",
		"student": updatedStudent,
	})
}

func (h *StudentHandler) DeleteStudent(c echo.Context) error {
	id := c.Param("id")

	if err := h.studentRepo.Delete(id); err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, echo.Map{
				"message": "Student not found",
			})
		}

		log.Errorf("Failed to delete student: %v", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to delete student",
		})
	}

	return c.JSON(http.StatusNoContent, nil)
}