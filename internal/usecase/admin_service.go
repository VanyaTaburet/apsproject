package usecase

import (
	"context"
	"errors"

	"github.com/vladimirfedunov/2chan-clone/internal/entity"
	"github.com/vladimirfedunov/2chan-clone/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type AdminService struct {
	repo repository.AdminRepository
}

func NewAdminService(repo repository.AdminRepository) *AdminService {
	return &AdminService{repo: repo}
}

func (s *AdminService) Login(ctx context.Context, username, password string) (*entity.Admin, error) {
	admin, err := s.repo.GetByUsername(ctx, username)
	if err != nil {
		return nil, errors.New("неверный логин или пароль")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(admin.PasswordHash()), []byte(password)); err != nil {
		return nil, errors.New("неверный логин или пароль")
	}
	_ = s.repo.UpdateLastLogin(ctx, admin.ID())
	return admin, nil
}

func (s *AdminService) CreateAdmin(ctx context.Context, username, password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	admin := entity.NewAdmin(username, string(hash))
	return s.repo.Create(ctx, admin)
}

func (s *AdminService) GetByID(ctx context.Context, id int64) (*entity.Admin, error) {
	// Реализуем через репозиторий (добавим метод в интерфейс и реализацию)
	return s.repo.GetByID(ctx, id)
}
