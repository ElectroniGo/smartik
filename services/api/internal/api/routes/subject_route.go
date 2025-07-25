package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/smartik/api/internal/api/handlers"
)

func RegisterSubjectRoutes(
	e *echo.Group,
	subjectHandler *handlers.SubjectHandler,
) {
	subjects := e.Group("/subjects")

	subjects.GET("", subjectHandler.GetAllSubjects).Name = "get_all_subjects"
	subjects.POST("/create", subjectHandler.CreateSubject).Name = "create_subject"
	subjects.GET("/:id", subjectHandler.GetSubjectById).Name = "get_subject_by_id"
	subjects.PATCH("/update/:id", subjectHandler.UpdateSubject).Name = "update_subject"
	subjects.DELETE("/delete/:id", subjectHandler.DeleteSubject).Name = "delete_subject"
}
