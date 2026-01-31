package app

import (
	"context"
	"html/template"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/vladimirfedunov/2chan-clone/internal/db/postgres"
	"github.com/vladimirfedunov/2chan-clone/internal/repository/postgresql"
	transport "github.com/vladimirfedunov/2chan-clone/internal/transport/http"
	"github.com/vladimirfedunov/2chan-clone/internal/usecase"
)

type App struct {
	router *gin.Engine
}

func NewApp() *App {
	// 1. Подключение к БД
	dsn := "postgres://postgres:110920@localhost:5432/chan?sslmode=disable"
	pool, err := postgres.NewPostgresPool(context.Background(), dsn)
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}

	// 2. Репозитории
	boardRepo := postgresql.NewPostgresBoardRepository(pool)
	threadRepo := postgresql.NewPostgresThreadRepository(pool)
	postRepo := postgresql.NewPostgresPostRepository(pool)
	adminRepo := postgresql.NewPostgresAdminRepository(pool)
	recentThreadRepo := postgresql.NewPostgresRecentThreadRepository(pool) // Передаем pool

	// 3. Сервисы (usecase)
	boardService := usecase.NewBoardService(boardRepo)
	threadService := usecase.NewThreadService(threadRepo, postRepo)
	postService := usecase.NewPostService(postRepo)
	adminService := usecase.NewAdminService(adminRepo)
	recentThreadService := usecase.NewRecentThreadService(recentThreadRepo) // Реальный репозиторий

	// 4. Handler
	handler := transport.NewHandler(boardService, threadService, postService, adminService, recentThreadService)

	// 5. Gin router в release mode
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	transport.SetupRoutes(router, handler, adminService)
	router.SetHTMLTemplate(loadTemplates())

	return &App{router: router}
}

func (a *App) Run() error {
	log.Println("2chan-clone server started on http://localhost:8080")
	return a.router.Run(":8080")
}

func loadTemplates() *template.Template {
	return template.Must(template.ParseGlob("web/templates/*.html"))
}
