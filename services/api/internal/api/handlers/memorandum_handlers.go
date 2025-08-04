package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/smartik/api/internal/service"
	"gorm.io/gorm"
)

type MemorandumHandler struct {
	service *service.MemorandumService
}

// Creates a new instance of MemorandumHandler
func NewMemorandumHandler(service *service.MemorandumService) *MemorandumHandler {
	return &MemorandumHandler{service}
}

// Handles the upload of a single memorandum file
func (h *MemorandumHandler) UploadMemorandum(c echo.Context) error {
	form, err := c.MultipartForm()
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid multipart form data",
			"error":   err.Error(),
		})
	}

	// Check if files are provided
	file := form.File["memorandum"]
	if len(file) == 0 {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "No memorandum file provided",
		})
	}

	examIdSlice, ok := form.Value["exam_id"]
	if !ok || len(examIdSlice) == 0 {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "No exam_id provided",
		})
	}

	examId := examIdSlice[0]

	result, err := h.service.UploadFile(file[0], examId, &service.MemorandumUploadResult{})
	if err != nil {
		log.Errorf("Upload service error: %v", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to upload memorandum",
		})
	}

	// Return partial content if some uploads failed
	if len(result.FailedUploads) > 0 {
		return c.JSON(http.StatusPartialContent, echo.Map{
			"message":            "Some memorandums failed to upload",
			"successful_uploads": len(result.SuccessfulUploads),
			"failed_uploads":     len(result.FailedUploads),
			"errors":             result.FailedUploads,
			"memorandums":        result.SuccessfulUploads,
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message":            "Memorandum uploaded successfully",
		"successful_uploads": len(result.SuccessfulUploads),
		"memorandums":        result.SuccessfulUploads,
	})
}

// Retrieves all memorandums
func (h *MemorandumHandler) GetAllMemorandums(c echo.Context) error {
	memorandums, err := h.service.GetAll()
	if err != nil {
		log.Errorf("Failed to get all memorandums: %v", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to retrieve memorandums",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message":     "Memorandums retrieved successfully",
		"memorandums": memorandums,
	})
}

// Retrieves a specific memorandum by ID
func (h *MemorandumHandler) GetMemorandumById(c echo.Context) error {
	id := c.Param("id")
	memorandum, err := h.service.GetById(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, echo.Map{
				"message": "Memorandum not found",
			})
		}

		log.Errorf("Failed to get memorandum by ID: %v", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to retrieve memorandum",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message":    "Memorandum retrieved successfully",
		"memorandum": memorandum,
	})
}

// Handles the serving of a memorandum file
func (h *MemorandumHandler) ServeMemorandumFile(c echo.Context) error {
	id := c.Param("id")

	// Get the file stream for the memorandum
	fileStream, err := h.service.GetFileStream(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, echo.Map{
				"message": "Memorandum file not found",
			})
		}
		log.Errorf("Failed to get file stream: %v", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to retrieve memorandum file",
		})
	}
	defer fileStream.Content.Close()

	// Set the response headers for file download
	c.Response().Header().Set(echo.HeaderContentType, fileStream.ContentType)
	c.Response().Header().Set(echo.HeaderContentLength, fmt.Sprintf("%d", fileStream.Size))
	c.Response().Header().Set(echo.HeaderContentDisposition, fmt.Sprintf("inline; filename=\"%s\"", fileStream.Filename))

	return c.Stream(http.StatusOK, fileStream.ContentType, fileStream.Content)
}

// Deletes a memorandum by ID
func (h *MemorandumHandler) DeleteMemorandum(c echo.Context) error {
	id := c.Param("id")
	if err := h.service.Delete(id); err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, echo.Map{
				"message": "Memorandum not found",
			})
		}

		log.Errorf("Failed to delete memorandum: %v", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to delete memorandum",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Memorandum deleted successfully",
	})
}
