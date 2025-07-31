package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/smartik/api/internal/models"
	"github.com/smartik/api/internal/service"
	"gorm.io/gorm"
)

// Handles HTTP requests for exam operations
type ExamHandler struct {
	service *service.ExamService
}

// Creates a new instance of ExamHandler
func NewExamHandler(service *service.ExamService) *ExamHandler {
	return &ExamHandler{service: service}
}

// Creates a new exam record
func (h *ExamHandler) CreateExam(c echo.Context) error {
	var exam models.Exam
	if err := c.Bind(&exam); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid input",
			"error":   err.Error(),
		})
	}

	if err := c.Validate(&exam); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Validation failed",
			"errors":  err.Error(),
		})
	}

	if err := h.service.Create(&exam); err != nil {
		log.Errorf("Failed to create exam: %v", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to create exam",
		})
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "Exam created successfully",
		"exam":    exam,
	})
}

// Retrieves all exams from the database
func (h *ExamHandler) GetAllExams(c echo.Context) error {
	exams, err := h.service.GetAll()
	if err != nil {
		log.Errorf("Failed to retrieve exams: %v", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to retrieve exams",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Exams retrieved successfully",
		"exams":   exams,
	})
}

// Retrieves a specific exam by ID
func (h *ExamHandler) GetExamById(c echo.Context) error {
	id := c.Param("id")
	exam, err := h.service.GetById(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, echo.Map{
				"message": "Exam not found",
			})
		}

		log.Errorf("Failed to retrieve exam: %v", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to retrieve exam",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Exam retrieved successfully",
		"exam":    exam,
	})
}

// Updates an existing exam record
func (h *ExamHandler) UpdateExam(c echo.Context) error {
	var updateData models.UpdateExam
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

	id := c.Param("id")
	updatedExam, err := h.service.Update(id, &updateData)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, echo.Map{
				"message": "Exam not found",
			})
		}

		log.Errorf("Failed to update exam: %v", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to update exam",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Exam updated successfully",
		"exam":    updatedExam,
	})
}

// Removes an exam from the database
func (h *ExamHandler) DeleteExam(c echo.Context) error {
	id := c.Param("id")
	if err := h.service.Delete(id); err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, echo.Map{
				"message": "Exam not found",
			})
		}

		log.Errorf("Failed to delete exam: %v", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to delete exam",
		})
	}

	return c.JSON(http.StatusNoContent, nil)
}
