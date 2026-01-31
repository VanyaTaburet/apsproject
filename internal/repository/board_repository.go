package repository

import (
	"context"

	"github.com/vladimirfedunov/2chan-clone/internal/entity"
)

type BoardRepository interface {
	GetAll(ctx context.Context) ([]*entity.Board, error)
	GetBySlug(ctx context.Context, slug string) (*entity.Board, error)
	Create(ctx context.Context, board *entity.Board) error
	Update(ctx context.Context, board *entity.Board) error
	Delete(ctx context.Context, slug string) error
}
