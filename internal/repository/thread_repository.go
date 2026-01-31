package repository

import (
	"context"

	"github.com/vladimirfedunov/2chan-clone/internal/entity"
)

type ThreadRepository interface {
	GetByBoard(ctx context.Context, boardSlug string) ([]*entity.Thread, error)
	GetByID(ctx context.Context, id int64) (*entity.Thread, error)
	Create(ctx context.Context, thread *entity.Thread) error
	Update(ctx context.Context, thread *entity.Thread) error
	Delete(ctx context.Context, id int64) error
	GetAllThreads(ctx context.Context) ([]*entity.Thread, error)
	GetRecent(ctx context.Context, boardSlug string, limit int) ([]*entity.Thread, error)
}
