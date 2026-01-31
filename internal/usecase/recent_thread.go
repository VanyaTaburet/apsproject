package usecase

import (
	"context"

	"github.com/vladimirfedunov/2chan-clone/internal/entity"
	"github.com/vladimirfedunov/2chan-clone/internal/repository"
)

type RecentThreadService interface {
	GetRecentThreads(ctx context.Context, limit int) ([]*entity.RecentThread, error)
	GetRecentThreadsByBoard(ctx context.Context, boardSlug string, limit int) ([]*entity.RecentThread, error)
}

type recentThreadService struct {
	repo repository.RecentThreadRepository
}

func NewRecentThreadService(repo repository.RecentThreadRepository) RecentThreadService {
	return &recentThreadService{repo: repo}
}

func (s *recentThreadService) GetRecentThreads(ctx context.Context, limit int) ([]*entity.RecentThread, error) {
	if limit <= 0 || limit > 50 {
		limit = 4 // показываем последние 4 треда
	}

	// Убираем все заглушки - только реальные данные из БД
	return s.repo.GetRecent(ctx, limit)
}

func (s *recentThreadService) GetRecentThreadsByBoard(ctx context.Context, boardSlug string, limit int) ([]*entity.RecentThread, error) {
	if limit <= 0 || limit > 50 {
		limit = 5
	}

	return s.repo.GetRecentByBoard(ctx, boardSlug, limit)
}
