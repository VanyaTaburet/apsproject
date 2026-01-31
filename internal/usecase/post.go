package usecase

import (
	"context"

	"github.com/vladimirfedunov/2chan-clone/internal/entity"
	"github.com/vladimirfedunov/2chan-clone/internal/repository"
)

type PostService interface {
	ListPosts(ctx context.Context, threadID int64, limit, offset int) ([]*entity.Post, error)
	GetPost(ctx context.Context, id int64) (*entity.Post, error)
	CreatePost(ctx context.Context, threadID int64, boardSlug, author, text string, imageURL *string, parentID *int64, tripcode *string, ipAddress string) (*entity.Post, error)
	DeletePost(ctx context.Context, id int64) error
}

type postService struct {
	repo repository.PostRepository
}

func NewPostService(repo repository.PostRepository) PostService {
	return &postService{repo: repo}
}

func (s *postService) ListPosts(ctx context.Context, threadID int64, limit, offset int) ([]*entity.Post, error) {
	return s.repo.GetRecent(ctx, threadID, limit) // offset можно добавить при необходимости
}

func (s *postService) GetPost(ctx context.Context, id int64) (*entity.Post, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *postService) CreatePost(ctx context.Context, threadID int64, boardSlug, author, text string, imageURL *string, parentID *int64, tripcode *string, ipAddress string) (*entity.Post, error) {
	post := entity.NewPost(threadID, boardSlug, author, text, imageURL, parentID, tripcode, ipAddress)
	if err := s.repo.Create(ctx, post); err != nil {
		return nil, err
	}
	return post, nil
}

func (s *postService) DeletePost(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}
