package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/smartik/api/internal/models"
	"github.com/smartik/api/internal/service"
	"gorm.io/gorm"
)

// Handles HTTP requests for answer script operations
type AnswerScriptHandler struct {
	service *service.AnswerScriptService
}

// Creates a new instance of AnswerScriptHandler
func NewAnswerScriptHandler(service *service.AnswerScriptService) *AnswerScriptHandler {
	return &AnswerScriptHandler{service: service}
}

// Handles the upload of multiple answer script files
func (h *AnswerScriptHandler) UploadScripts(c echo.Context) error {
	form, err := c.MultipartForm()
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid multipart form data",
			"error":   err.Error(),
		})
	}

	// Check if files are provided
	files := form.File["answer_scripts"]
	if len(files) == 0 {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "No answer scripts provided",
		})
	}

	result, err := h.service.UploadFiles(files)
	if err != nil {
		log.Errorf("Upload service error: %v", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Upload service error",
		})
	}

	// Return partial content if some uploads failed
	if len(result.FailedUploads) > 0 {
		return c.JSON(http.StatusPartialContent, echo.Map{
			"message":            "Some answer scripts failed to upload",
			"successful_uploads": len(result.SuccessfulUploads),
			"failed_uploads":     len(result.FailedUploads),
			"errors":             result.FailedUploads,
			"answer_scripts":     result.SuccessfulUploads,
		})
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message":        "Answer scripts uploaded successfully",
		"count":          len(result.SuccessfulUploads),
		"answer_scripts": result.SuccessfulUploads,
	})
}

// Retrieves all answer scripts from the database
func (h *AnswerScriptHandler) GetAllScripts(c echo.Context) error {
	answerScripts, err := h.service.GetAll()
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

// Retrieves a specific answer script by ID
func (h *AnswerScriptHandler) GetScriptById(c echo.Context) error {
	id := c.Param("id")
	answerScript, err := h.service.GetById(id)
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

// Serves the actual file content for download or viewing
func (h *AnswerScriptHandler) ServeAnswerScript(c echo.Context) error {
	id := c.Param("id")

	// Get file stream
	fileStream, err := h.service.GetFileStream(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, echo.Map{
				"message": "Answer script not found",
			})
		}
		log.Errorf("Failed to get file stream: %v", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to retrieve answer script file",
		})
	}
	defer fileStream.Content.Close()

	// Set response headers for file serving
	c.Response().Header().Set(echo.HeaderContentType, fileStream.ContentType)
	c.Response().Header().Set(echo.HeaderContentLength, fmt.Sprintf("%d", fileStream.Size))
	c.Response().Header().Set(echo.HeaderContentDisposition,
		fmt.Sprintf("inline; filename=\"%s\"", fileStream.Filename))

	return c.Stream(http.StatusOK, fileStream.ContentType, fileStream.Content)
}

// Updates an existing answer script record
func (h *AnswerScriptHandler) UpdateScript(c echo.Context) error {
	id := c.Param("id")

	var updateData models.AnswerScript
	if err := c.Bind(&updateData); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid input",
			"error":   err.Error(),
		})
	}

	updatedScript, err := h.service.Update(id, &updateData)
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

// Removes an answer script from both database and storage
func (h *AnswerScriptHandler) DeleteScript(c echo.Context) error {
	id := c.Param("id")

	if err := h.service.Delete(id); err != nil {
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
