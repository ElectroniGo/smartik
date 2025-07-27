package handlers

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	minio "github.com/minio/minio-go/v7"
	"github.com/smartik/api/internal/config"
	"github.com/smartik/api/internal/models"
	"github.com/smartik/api/internal/repository"
	"gorm.io/gorm"
)

type AnswerScriptHandler struct {
	repo        *repository.AnswerScriptRepository
	minioClient *minio.Client
	minioBucket string
}

func NewAnswerScriptHandler(repo *repository.AnswerScriptRepository,
	minioClient *minio.Client,
) (*AnswerScriptHandler, error) {

	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}

	return &AnswerScriptHandler{repo, minioClient, cfg.MinioStorageBucket}, nil
}

func (h *AnswerScriptHandler) UploadScripts(c echo.Context) error {
	data, err := c.MultipartForm()
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid multipart form data",
			"error":   err.Error(),
		})
	}

	if len(data.File["answer_scripts"]) < 1 {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "No answer scripts provided",
		})
	}

	errors := map[string]any{
		"count": 0,
	}

	for _, file := range data.File["answer_scripts"] {
		src, err := file.Open()
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": "Failed to read file",
				"file":    file.Filename,
				"error":   err.Error(),
			})
		}
		defer src.Close()

		// upload to MinIO
		info, err := h.minioClient.PutObject(context.Background(), h.minioBucket,
			file.Filename, src, file.Size, minio.PutObjectOptions{
				ContentType: file.Header.Get("Content-Type"),
			})
		if err != nil {
			// Saves error details for each failed upload & moves to next file
			errors["file_name"] = map[string]any{
				"filename": file.Filename,
				"error":    err.Error(),
			}
			errors["count"] = errors["count"].(int) + 1
			continue
		}

		// Create answer script record in the database
		answerScript := &models.AnswerScript{
			FileName: file.Filename,
			FileUrl:  &info.Location,
			Status:   models.StatusUploaded,
		}

		if err := h.repo.Create(answerScript); err != nil {
			log.Errorf("Failed to save answer script record: %v", err)
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"message": "Failed to save answer script record",
			})
		}
	}

	if errors["count"].(int) > 0 {
		return c.JSON(http.StatusPartialContent, echo.Map{
			"message": "Some answer scripts failed to upload",
			"errors":  errors,
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Answer scripts uploaded successfully",
		"count":   len(data.File["answer_scripts"]),
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
