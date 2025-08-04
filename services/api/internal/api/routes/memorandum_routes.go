package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/smartik/api/internal/api/handlers"
)

func RegisterMemorandumRoutes(e *echo.Group, handlers *handlers.MemorandumHandler) {
	memorandums := e.Group("/memorandums")

	memorandums.POST("/upload", handlers.UploadMemorandum).Name = "upload_memorandum"
	memorandums.GET("", handlers.GetAllMemorandums).Name = "get_all_memorandums"
	memorandums.GET("/:id", handlers.GetMemorandumById).Name = "get_memorandum_by_id"
	memorandums.GET("/serve/:id", handlers.ServeMemorandumFile).Name = "serve_memorandum_file"
	// TODO: Implement update functionality
	memorandums.DELETE("/delete/:id", handlers.DeleteMemorandum).Name = "delete_memorandum"
}
