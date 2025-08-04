package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/smartik/api/internal/api/handlers"
	"github.com/smartik/api/internal/api/routes"
	"github.com/smartik/api/internal/config"
	"github.com/smartik/api/internal/models"
	"github.com/smartik/api/internal/repository"
	"github.com/smartik/api/internal/repository/minio"
	"github.com/smartik/api/internal/repository/postgres"
	"github.com/smartik/api/internal/service"
)

var startTime time.Time

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Warnf("Failed to load config: %v (Using defaults)", err)
	}

	// Database connection
	db, err := postgres.NewConnection(cfg.PostgresURI)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			log.Fatalf("Failed to get database instance: %v", err)
		}
		if err := sqlDB.Close(); err != nil {
			log.Fatalf("Failed to close database connection: %v", err)
		}
	}()

	if err := db.AutoMigrate(models.GetAllModels()...); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Initialize MinIO client
	minioClient, err := minio.NewMinioClient(cfg.MinioEndpointUrl,
		cfg.MinioAccessId, cfg.MinioSecretKey, cfg,
	)
	if err != nil {
		log.Fatalf("Something went wrong: %v", err)
	}

	// Initialize repositories and handlers
	studentRepo := repository.NewStudentRepository(db)
	subjectRepo := repository.NewSubjectRepository(db)
	examRepo := repository.NewExamRepository(db)
	answerScriptRepo := repository.NewAnswerScriptRepository(db)
	memorandumRepo := repository.NewMemorandumRepository(db)

	// Initialize services
	studentService := service.NewStudentService(studentRepo)
	subjectService := service.NewSubjectService(subjectRepo)
	examService := service.NewExamService(examRepo)
	answerScriptService := service.NewAnswerScriptService(answerScriptRepo, minioClient, cfg)
	memorandumService := service.NewMemorandumService(memorandumRepo, minioClient, cfg)

	// Initialize handlers
	studentHandler := handlers.NewStudentHandler(studentService)
	subjectHandler := handlers.NewSubjectHandler(subjectService)
	examHandler := handlers.NewExamHandler(examService)
	answerScriptHandler := handlers.NewAnswerScriptHandler(answerScriptService)
	memorandumHandler := handlers.NewMemorandumHandler(memorandumService)

	// Create Echo instance
	e := echo.New()
	e.Validator = NewCustomValidator()
	addr := fmt.Sprintf(":%s", cfg.Port)
	startTime = time.Now() // Record the start time

	// Middlware
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		LogLevel: log.ERROR,
	}))

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[${time_rfc3339}] ${remote_ip} ${method} ${path} ${status}\n",
	}))

	// Routes
	v1 := e.Group("/api/v1")
	{
		v1.Any("/health", func(c echo.Context) error {
			t := time.Since(startTime)

			return c.JSON(http.StatusOK, echo.Map{
				"status": "healthy",
				"time":   time.Now().Format(time.RFC3339),
				"uptime": string(fmt.Sprintf("%d Hours, %d Minutes, %d Seconds",
					int(t.Abs().Hours()),
					int(t.Abs().Minutes()),
					int(t.Abs().Seconds()),
				)),
			})
		})

		v1.GET("/reference", func(c echo.Context) error {
			return c.JSON(http.StatusOK, echo.Map{
				"message": "Available Routes Reference",
				"routes":  e.Routes(),
			})
		})

		routes.RegisterStudentRoutes(v1, studentHandler)
		routes.RegisterSubjectRoutes(v1, subjectHandler)
		routes.RegisterExamRoutes(v1, examHandler)
		routes.RegisterAnswerScriptRoutes(v1, answerScriptHandler)
		routes.RegisterMemorandumRoutes(v1, memorandumHandler)
	}

	go func() {
		if err := e.Start(addr); err != nil {
			e.Logger.Infof("Shutting down server: %v", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatalf("Failed to gracefully shutdown server: %v", err)
	}
}
