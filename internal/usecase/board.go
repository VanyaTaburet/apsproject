package usecase

import (
	"context"

	"github.com/vladimirfedunov/2chan-clone/internal/entity"
	"github.com/vladimirfedunov/2chan-clone/internal/repository"
)

type BoardService interface {
	ListBoards(ctx context.Context) ([]*entity.Board, error)
	GetBoard(ctx context.Context, slug string) (*entity.Board, error)
	CreateBoard(ctx context.Context, slug, name, description string) error
	UpdateBoard(ctx context.Context, slug, name, description string) error
	DeleteBoard(ctx context.Context, slug string) error
}

type boardService struct {
	repo repository.BoardRepository
}

func NewBoardService(repo repository.BoardRepository) BoardService {
	return &boardService{repo: repo}
}

func (s *boardService) ListBoards(ctx context.Context) ([]*entity.Board, error) {
	return s.repo.GetAll(ctx)
}

func (s *boardService) GetBoard(ctx context.Context, slug string) (*entity.Board, error) {
	return s.repo.GetBySlug(ctx, slug)
}

func (s *boardService) CreateBoard(ctx context.Context, slug, name, description string) error {
	board := entity.NewBoard(slug, name, description)
	return s.repo.Create(ctx, board)
}

func (s *boardService) UpdateBoard(ctx context.Context, slug, name, description string) error {
	board := entity.NewBoard(slug, name, description)
	return s.repo.Update(ctx, board)
}

func (s *boardService) DeleteBoard(ctx context.Context, slug string) error {
	return s.repo.Delete(ctx, slug)
}
