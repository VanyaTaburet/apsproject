package usecase

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/vladimirfedunov/2chan-clone/internal/entity"
)

type mockBoardRepository struct {
	mock.Mock
}

func (m *mockBoardRepository) GetAll(ctx context.Context) ([]*entity.Board, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*entity.Board), args.Error(1)
}

func (m *mockBoardRepository) GetBySlug(ctx context.Context, slug string) (*entity.Board, error) {
	args := m.Called(ctx, slug)
	return args.Get(0).(*entity.Board), args.Error(1)
}

func (m *mockBoardRepository) Create(ctx context.Context, board *entity.Board) error {
	args := m.Called(ctx, board)
	return args.Error(0)
}

func (m *mockBoardRepository) Update(ctx context.Context, board *entity.Board) error {
	args := m.Called(ctx, board)
	return args.Error(0)
}

func (m *mockBoardRepository) Delete(ctx context.Context, slug string) error {
	args := m.Called(ctx, slug)
	return args.Error(0)
}

/*
   Tests
*/

func TestBoardService_ListBoards(t *testing.T) {
	ctx := context.Background()

	repo := new(mockBoardRepository)
	service := NewBoardService(repo)

	boards := []*entity.Board{
		entity.NewBoard("b", "Random", "Random board"),
	}

	repo.
		On("GetAll", ctx).
		Return(boards, nil)

	result, err := service.ListBoards(ctx)

	assert.NoError(t, err)
	assert.Equal(t, boards, result)
	repo.AssertExpectations(t)
}

func TestBoardService_GetBoard(t *testing.T) {
	ctx := context.Background()

	repo := new(mockBoardRepository)
	service := NewBoardService(repo)

	board := entity.NewBoard("tech", "Technology", "Tech board")

	repo.
		On("GetBySlug", ctx, "tech").
		Return(board, nil)

	result, err := service.GetBoard(ctx, "tech")

	assert.NoError(t, err)
	assert.Equal(t, board, result)
	repo.AssertExpectations(t)
}

func TestBoardService_CreateBoard(t *testing.T) {
	ctx := context.Background()

	repo := new(mockBoardRepository)
	service := NewBoardService(repo)

	repo.
		On("Create", ctx, mock.MatchedBy(func(b *entity.Board) bool {
			return b.Slug == "news" &&
				b.Name == "News" &&
				b.Description == "News board"
		})).
		Return(nil)

	err := service.CreateBoard(ctx, "news", "News", "News board")

	assert.NoError(t, err)
	repo.AssertExpectations(t)
}

func TestBoardService_UpdateBoard(t *testing.T) {
	ctx := context.Background()

	repo := new(mockBoardRepository)
	service := NewBoardService(repo)

	repo.
		On("Update", ctx, mock.MatchedBy(func(b *entity.Board) bool {
			return b.Slug == "games" &&
				b.Name == "Games" &&
				b.Description == "Games board"
		})).
		Return(nil)

	err := service.UpdateBoard(ctx, "games", "Games", "Games board")

	assert.NoError(t, err)
	repo.AssertExpectations(t)
}

func TestBoardService_DeleteBoard(t *testing.T) {
	ctx := context.Background()

	repo := new(mockBoardRepository)
	service := NewBoardService(repo)

	repo.
		On("Delete", ctx, "b").
		Return(nil)

	err := service.DeleteBoard(ctx, "b")

	assert.NoError(t, err)
	repo.AssertExpectations(t)
}
