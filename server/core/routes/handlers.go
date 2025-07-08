package routes

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	textextractor "github.com/smartik/core/internal/text_extractor"
)

type TriggerRequest struct {
	PdfUrl string `json:"url"`
}

func onTrigger(c echo.Context) error {
	req := c.Request().Body
	defer req.Close()

	buf, err := io.ReadAll(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"success": false,
			"detail":  "Failed to read request body",
		})
	}

	var reqData TriggerRequest
	if err := json.Unmarshal(buf, &reqData); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"success": false,
			"detail":  "Invalid request format",
		})
	}

	pdfBuf, err := textextractor.ReadScript(reqData.PdfUrl)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"success": false,
			"detail":  "Failed to fetch PDF",
		})
	}

	scriptText, err := textextractor.ExtractText(pdfBuf)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"success": false,
			"detail":  "Failed to extract text from PDF",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"success": true,
		"detail":  "Triggered successfully",
		"data":    scriptText,
	})
}
