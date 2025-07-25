package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/smartik/api/internal/models"
	"github.com/smartik/api/internal/repository"
	"gorm.io/gorm"
)

type ExamHandler struct {
	ExamRepo *repository.ExamRepository
}

func NewExamHandler(repo *repository.ExamRepository) *ExamHandler {
	return &ExamHandler{repo}
}

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

	if err := h.ExamRepo.Create(&exam); err != nil {
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

func (h *ExamHandler) GetAllExams(c echo.Context) error {
	exams, err := h.ExamRepo.GetAll()
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

func (h *ExamHandler) GetExamById(c echo.Context) error {
	id := c.Param("id")
	exam, err := h.ExamRepo.GetById(id)
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
	updatedExam, err := h.ExamRepo.Update(id, &updateData)
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

func (h *ExamHandler) DeleteExam(c echo.Context) error {
	id := c.Param("id")
	if err := h.ExamRepo.Delete(id); err != nil {
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
