package routes

import "github.com/labstack/echo/v4"

func Attatch(e *echo.Echo) {
	e.POST("/trigger", onTrigger)
}
