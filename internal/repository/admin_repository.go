package repository

import (
	"context"

	"github.com/vladimirfedunov/2chan-clone/internal/entity"
)

type AdminRepository interface {
	GetByUsername(ctx context.Context, username string) (*entity.Admin, error)
	Create(ctx context.Context, admin *entity.Admin) error
	UpdateLastLogin(ctx context.Context, id int64) error
	GetByID(ctx context.Context, id int64) (*entity.Admin, error)
}
