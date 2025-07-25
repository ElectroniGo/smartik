package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/smartik/api/internal/models"
	"github.com/smartik/api/internal/repository"
	"gorm.io/gorm"
)

type SubjectHandler struct {
	subjectRepo *repository.SubjectRepository
}

func NewSubjectHandler(repo *repository.SubjectRepository) *SubjectHandler {
	return &SubjectHandler{repo}
}

func (h *SubjectHandler) CreateSubject(c echo.Context) error {
	var subject models.Subject

	if err := c.Bind(&subject); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid input",
			"error":   err.Error(),
		})
	}

	if err := c.Validate(&subject); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Validation failed",
			"errors":  err.Error(),
		})
	}

	if err := h.subjectRepo.Create(&subject); err != nil {
		log.Errorf("Failed to create subject: %v", err)

		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to create subject",
		})
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "Subject created successfully",
		"subject": subject,
	})
}

func (h *SubjectHandler) GetAllSubjects(c echo.Context) error {
	subjects, err := h.subjectRepo.GetAll()
	if err != nil {
		log.Errorf("Failed to retrieve subjects: %v", err)

		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to retrieve subjects",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message":  "Subjects retrieved successfully",
		"subjects": subjects,
	})
}

func (h *SubjectHandler) GetSubjectById(c echo.Context) error {
	id := c.Param("id")
	subject, err := h.subjectRepo.GetById(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, echo.Map{
				"message": "Subject not found",
			})
		}

		log.Errorf("Failed to retrieve subject by ID %s: %v", id, err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to retrieve subject",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Subject retrieved successfully",
		"subject": subject,
	})
}

func (h *SubjectHandler) UpdateSubject(c echo.Context) error {
	id := c.Param("id")
	var updateData models.UpdateSubject

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

	updatedSubject, err := h.subjectRepo.Update(id, &updateData)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, echo.Map{
				"message": "Subject not found",
			})
		}

		log.Errorf("Failed to update subject: %v", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to update subject",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Subject updated successfully",
		"subject": updatedSubject,
	})
}

func (h *SubjectHandler) DeleteSubject(c echo.Context) error {
	id := c.Param("id")
	if err := h.subjectRepo.Delete(id); err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, echo.Map{
				"message": "Subject not found",
			})
		}

		log.Errorf("Failed to delete subject: %v", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to delete subject",
		})
	}

	return c.JSON(http.StatusNoContent, nil)
}
