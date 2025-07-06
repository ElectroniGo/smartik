package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func onTrigger(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"success": true,
		"detail":  "Triggered successfully",
	})
}
