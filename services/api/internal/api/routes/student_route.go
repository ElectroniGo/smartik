package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/smartik/api/internal/api/handlers"
)

func RegisterStudentRoutes(e *echo.Group, studentHandler *handlers.StudentHandler) {
	students := e.Group("/students")

	students.GET("", studentHandler.GetAllStudents).Name = "get_all_students"
	students.POST("/create", studentHandler.CreateStudent).Name = "create_student"
	students.GET("/:id", studentHandler.GetStudentById).Name = "get_student_by_exam_number"
	students.PATCH("/update/:id", studentHandler.UpdateStudent).Name = "update_student"
	students.DELETE("/delete/:id", studentHandler.DeleteStudent).Name = "delete_student"
}
