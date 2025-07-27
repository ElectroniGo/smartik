package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/smartik/api/internal/models"
	"github.com/smartik/api/internal/repository"
	"gorm.io/gorm"
)

type AnswerScriptHandler struct {
	repo *repository.AnswerScriptRepository
}

func NewAnswerScriptHandler(repo *repository.AnswerScriptRepository,
) *AnswerScriptHandler {
	return &AnswerScriptHandler{repo}
}

func (h *AnswerScriptHandler) UploadScripts(c echo.Context) error {
	var answerScripts []models.AnswerScript
	if err := c.Bind(&answerScripts); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid input",
			"error":   err.Error(),
		})
	}

	for i, script := range answerScripts {
		if err := c.Validate(&script); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": "Validation error",
				"index":   i, // Index of the script that failed validation
				"errors":  err.Error(),
			})
		}
	}

	if err := h.repo.Create(&answerScripts); err != nil {
		log.Errorf("Failed to create answer scripts: %v", err)

		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to create answer scripts",
		})
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message":        "Answer scripts created successfully",
		"count":          len(answerScripts),
		"answer_scripts": answerScripts,
	})
}

func (h *AnswerScriptHandler) GetAllScripts(c echo.Context) error {
	answerScripts, err := h.repo.GetAll()
	if err != nil {
		log.Errorf("Failed to get all answer scripts: %v", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to retrieve answer scripts",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message":        "Answer scripts retrieved successfully",
		"answer_scripts": answerScripts,
	})
}

func (h *AnswerScriptHandler) GetScriptById(c echo.Context) error {
	id := c.Param("id")
	answerScript, err := h.repo.GetById(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, echo.Map{
				"message": "Answer script not found",
			})
		}

		log.Errorf("Failed to get answer script by ID: %v", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to retrieve answer script",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message":       "Answer script retrieved successfully",
		"answer_script": answerScript,
	})
}

func (h *AnswerScriptHandler) UpdateScript(c echo.Context) error {
	id := c.Param("id")

	var updateData models.AnswerScript
	if err := c.Bind(&updateData); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid input",
			"error":   err.Error(),
		})
	}

	updatedScript, err := h.repo.Update(id, &updateData)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, echo.Map{
				"message": "Answer script not found",
			})
		}

		log.Errorf("Failed to update answer script: %v", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to update answer script",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message":       "Answer script updated successfully",
		"answer_script": updatedScript,
	})
}

func (h *AnswerScriptHandler) DeleteScript(c echo.Context) error {
	id := c.Param("id")
	if err := h.repo.Delete(id); err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, echo.Map{
				"message": "Answer script not found",
			})
		}

		log.Errorf("Failed to delete answer script: %v", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to delete answer script",
		})
	}

	return c.JSON(http.StatusNoContent, nil)
}
