package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/smartik/api/internal/api/handlers"
)

func RegisterExamRoutes(e *echo.Group, handler *handlers.ExamHandler) {
	exams := e.Group("/exams")

	exams.GET("", handler.GetAllExams).Name = "get_all_exams"
	exams.POST("/create", handler.CreateExam).Name = "create_exam"
	exams.GET("/:id", handler.GetExamById).Name = "get_exam_by_id"
	exams.PATCH("/update/:id", handler.UpdateExam).Name = "update_exam"
	exams.DELETE("/delete/:id", handler.DeleteExam).Name = "delete_exam"
}
