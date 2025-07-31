package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/smartik/api/internal/models"
	"github.com/smartik/api/internal/service"
	"gorm.io/gorm"
)

// Handles HTTP requests for student operations
type StudentHandler struct {
	service *service.StudentService
}

// Creates a new instance of StudentHandler
func NewStudentHandler(service *service.StudentService) *StudentHandler {
	return &StudentHandler{service: service}
}

// Creates a new student record
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

	if err := h.service.Create(&newStudent); err != nil {
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

// Retrieves all students from the database
func (h *StudentHandler) GetAllStudents(c echo.Context) error {
	students, err := h.service.GetAll()
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

// Retrieves a specific student by ID
func (h *StudentHandler) GetStudentById(c echo.Context) error {
	id := c.Param("id")

	student, err := h.service.GetById(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, echo.Map{
				"message": "Student not found",
			})
		}

		log.Errorf("Failed to get student by ID: %v", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to retrieve student",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Student retrieved successfully",
		"student": student,
	})
}

// Updates an existing student record
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

	updatedStudent, err := h.service.Update(id, &updateData)
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

// Removes a student from the database
func (h *StudentHandler) DeleteStudent(c echo.Context) error {
	id := c.Param("id")

	if err := h.service.Delete(id); err != nil {
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
