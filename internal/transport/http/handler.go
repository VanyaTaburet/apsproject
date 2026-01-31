package http

import (
	"github.com/vladimirfedunov/2chan-clone/internal/usecase"
)

type Handler struct {
	BoardService        usecase.BoardService
	ThreadService       usecase.ThreadService
	PostService         usecase.PostService
	AdminService        *usecase.AdminService // Указатель на структуру
	RecentThreadService usecase.RecentThreadService
}

func NewHandler(
	boardService usecase.BoardService,
	threadService usecase.ThreadService,
	postService usecase.PostService,
	adminService *usecase.AdminService, // Изменить на указатель
	recentThreadService ...usecase.RecentThreadService,
) *Handler {
	h := &Handler{
		BoardService:  boardService,
		ThreadService: threadService,
		PostService:   postService,
		AdminService:  adminService,
	}

	if len(recentThreadService) > 0 {
		h.RecentThreadService = recentThreadService[0]
	}

	return h
}
