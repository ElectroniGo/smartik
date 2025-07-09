package routes

import "github.com/labstack/echo/v4"

func Attach(e *echo.Echo) {
	e.POST("/trigger", onTrigger)
}
