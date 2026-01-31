package usecase

import (
	"context"

	"github.com/vladimirfedunov/2chan-clone/internal/entity"
	"github.com/vladimirfedunov/2chan-clone/internal/repository"
)

type ThreadService interface {
	ListThreads(ctx context.Context, boardSlug string, limit, offset int) ([]*entity.Thread, error)
	GetThread(ctx context.Context, id int64) (*entity.Thread, []*entity.Post, error)
	CreateThread(ctx context.Context, boardSlug, author, title, text string, imageURL *string, tripcode *string, ipAddress string) (*entity.Thread, error)
	StickyThread(ctx context.Context, id int64, sticky bool) error
	LockThread(ctx context.Context, id int64, locked bool) error
	DeleteThread(ctx context.Context, id int64) error
	GetAllThreads(ctx context.Context) ([]*entity.Thread, error) // <-- add this line
	ListThreadPreviews(ctx context.Context, boardSlug string) ([]ThreadPreview, error)
}

type threadService struct {
	threadRepo repository.ThreadRepository
	postRepo   repository.PostRepository
}

func NewThreadService(threadRepo repository.ThreadRepository, postRepo repository.PostRepository) ThreadService {
	return &threadService{threadRepo: threadRepo, postRepo: postRepo}
}

func (s *threadService) ListThreads(ctx context.Context, boardSlug string, limit, offset int) ([]*entity.Thread, error) {
	return s.threadRepo.GetRecent(ctx, boardSlug, limit) // offset можно добавить при необходимости
}

func (s *threadService) GetThread(ctx context.Context, id int64) (*entity.Thread, []*entity.Post, error) {
	thread, err := s.threadRepo.GetByID(ctx, id)
	if err != nil {
		return nil, nil, err
	}
	posts, err := s.postRepo.GetByThread(ctx, id)
	if err != nil {
		return nil, nil, err
	}
	return thread, posts, nil
}

func (s *threadService) CreateThread(ctx context.Context, boardSlug, author, title, text string, imageURL *string, tripcode *string, ipAddress string) (*entity.Thread, error) {
	thread := entity.NewThread(boardSlug, title, author, false, false)
	if err := s.threadRepo.Create(ctx, thread); err != nil {
		return nil, err
	}
	post := entity.NewPost(thread.ID(), boardSlug, author, text, imageURL, nil, tripcode, ipAddress)
	if err := s.postRepo.Create(ctx, post); err != nil {
		return nil, err
	}
	return thread, nil
}

func (s *threadService) StickyThread(ctx context.Context, id int64, sticky bool) error {
	thread, err := s.threadRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	thread.SetSticky(sticky)
	return s.threadRepo.Update(ctx, thread)
}

func (s *threadService) LockThread(ctx context.Context, id int64, locked bool) error {
	thread, err := s.threadRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	thread.SetLocked(locked)
	return s.threadRepo.Update(ctx, thread)
}

func (s *threadService) DeleteThread(ctx context.Context, id int64) error {
	return s.threadRepo.Delete(ctx, id)
}

func (s *threadService) GetAllThreads(ctx context.Context) ([]*entity.Thread, error) {
	return s.threadRepo.GetAllThreads(ctx)
}

type ThreadPreview struct {
	Thread *entity.Thread
	Posts  []*entity.Post // OP + до 3 последних комментариев
}

func (s *threadService) ListThreadPreviews(ctx context.Context, boardSlug string) ([]ThreadPreview, error) {
	threads, err := s.threadRepo.GetByBoard(ctx, boardSlug)
	if err != nil {
		return nil, err
	}
	previews := make([]ThreadPreview, 0, len(threads))
	for _, t := range threads {
		posts, err := s.postRepo.GetByThread(ctx, t.ID())
		if err != nil || len(posts) == 0 {
			previews = append(previews, ThreadPreview{Thread: t, Posts: nil})
			continue
		}
		// Гарантируем, что у Thread установлен ID
		t.SetID(posts[0].ThreadID())
		var showPosts []*entity.Post
		if len(posts) > 4 {
			showPosts = append(showPosts, posts[0])                // OP
			showPosts = append(showPosts, posts[len(posts)-3:]...) // последние 3
		} else {
			showPosts = posts
		}
		previews = append(previews, ThreadPreview{Thread: t, Posts: showPosts})
	}
	return previews, nil
}
