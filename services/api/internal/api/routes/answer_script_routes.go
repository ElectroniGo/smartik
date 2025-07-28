package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/smartik/api/internal/api/handlers"
)

func RegisterAnswerScriptRoutes(e *echo.Group, handlers *handlers.AnswerScriptHandler) {
	answerScripts := e.Group("/scripts")

	answerScripts.POST("/upload", handlers.UploadScripts).Name = "upload_answer_scripts"
	answerScripts.GET("", handlers.GetAllScripts).Name = "get_all_answer_scripts"
	answerScripts.GET("/:id", handlers.GetScriptById).Name = "get_answer_script_by_id"
	answerScripts.GET("/serve/:id", handlers.ServeAnswerScript).Name = "serve_answer_script_file"
	answerScripts.PATCH("/update/:id", handlers.UpdateScript).Name = "update_answer_script"
	answerScripts.DELETE("/delete/:id", handlers.DeleteScript).Name = "delete_answer_script"
}
