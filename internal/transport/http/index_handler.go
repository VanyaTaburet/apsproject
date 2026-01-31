package http

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vladimirfedunov/2chan-clone/internal/entity"
	"github.com/vladimirfedunov/2chan-clone/internal/usecase"
)

type IndexHandler struct {
	recentThreadService usecase.RecentThreadService
	boardService        usecase.BoardService
}

func NewIndexHandler(
	recentThreadService usecase.RecentThreadService,
	boardService usecase.BoardService,
) *IndexHandler {
	return &IndexHandler{
		recentThreadService: recentThreadService,
		boardService:        boardService,
	}
}

func (h *Handler) IndexPage(c *gin.Context) {
	// Получаем список борд
	boards, err := h.BoardService.ListBoards(c.Request.Context())
	if err != nil {
		c.String(http.StatusInternalServerError, "Ошибка: %v", err)
		return
	}

	// Получаем последние треды (если сервис доступен)
	var recentThreads []*entity.RecentThread
	if h.RecentThreadService != nil {
		recentThreads, err = h.RecentThreadService.GetRecentThreads(c.Request.Context(), 10)
		if err != nil {
			// Логируем ошибку, но не прерываем загрузку
			log.Printf("Error loading recent threads: %v", err)
			recentThreads = []*entity.RecentThread{}
		}
	}

	// Для отладки - проверим, что данные есть
	log.Printf("Boards count: %d, Recent threads count: %d", len(boards), len(recentThreads))

	c.HTML(http.StatusOK, "index.html", gin.H{
		"Boards":        boards,
		"RecentThreads": recentThreads,
	})
}

// Добавляем новый метод для страницы "О проекте"
func (h *Handler) AboutPage(c *gin.Context) {
	ctx := context.Background()

	// Получаем статистику для отображения
	boards, err := h.BoardService.ListBoards(ctx)
	if err != nil {
		log.Printf("Error getting boards for about page: %v", err)
	}

	// Здесь можно добавить методы для получения статистики тредов и постов
	// Пока используем заглушки
	data := gin.H{
		"BoardsCount":  len(boards),
		"ThreadsCount": "4+",  // Можно сделать реальный подсчет
		"PostsCount":   "10+", // Можно сделать реальный подсчет
	}

	c.HTML(http.StatusOK, "about.html", data)
}
