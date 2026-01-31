package repository

import (
	"context"

	"github.com/vladimirfedunov/2chan-clone/internal/entity"
)

type PostRepository interface {
	GetByThread(ctx context.Context, threadID int64) ([]*entity.Post, error)
	GetByID(ctx context.Context, id int64) (*entity.Post, error)
	Create(ctx context.Context, post *entity.Post) error
	Delete(ctx context.Context, id int64) error
	GetRecent(ctx context.Context, threadID int64, limit int) ([]*entity.Post, error)
}
