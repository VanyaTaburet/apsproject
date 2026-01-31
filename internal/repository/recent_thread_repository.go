package repository

import (
	"context"

	"github.com/vladimirfedunov/2chan-clone/internal/entity"
)

type RecentThreadRepository interface {
	GetRecent(ctx context.Context, limit int) ([]*entity.RecentThread, error)
	GetRecentByBoard(ctx context.Context, boardSlug string, limit int) ([]*entity.RecentThread, error)
}
