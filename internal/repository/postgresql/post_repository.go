package postgresql

import (
	"context"

	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vladimirfedunov/2chan-clone/internal/entity"
	"github.com/vladimirfedunov/2chan-clone/internal/repository"
)

type PostgresPostRepository struct {
	pool *pgxpool.Pool
}

func NewPostgresPostRepository(pool *pgxpool.Pool) repository.PostRepository {
	return &PostgresPostRepository{pool: pool}
}

func (r *PostgresPostRepository) GetByThread(ctx context.Context, threadID int64) ([]*entity.Post, error) {
	rows, err := r.pool.Query(ctx, `SELECT id, thread_id, board_slug, author, text, created_at, image_url, parent_id, tripcode, ip_address FROM posts WHERE thread_id=$1 ORDER BY created_at ASC`, threadID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var posts []*entity.Post
	for rows.Next() {
		var id int64
		var threadID int64
		var boardSlug, author, text string
		var createdAt time.Time
		var imageURL *string
		var parentID *int64
		var tripcode *string
		var ipAddress string
		if err := rows.Scan(&id, &threadID, &boardSlug, &author, &text, &createdAt, &imageURL, &parentID, &tripcode, &ipAddress); err != nil {
			return nil, err
		}
		post := entity.NewPost(threadID, boardSlug, author, text, imageURL, parentID, tripcode, ipAddress)
		post.SetID(id)
		post.SetCreatedAt(createdAt)
		posts = append(posts, post)
	}
	return posts, nil
}

func (r *PostgresPostRepository) GetByID(ctx context.Context, id int64) (*entity.Post, error) {
	row := r.pool.QueryRow(ctx, `SELECT id, thread_id, board_slug, author, text, created_at, image_url, parent_id, tripcode, ip_address FROM posts WHERE id=$1`, id)
	var (
		pid                     int64
		threadID                int64
		boardSlug, author, text string
		createdAt               time.Time
		imageURL                *string
		parentID                *int64
		tripcode                *string
		ipAddress               string
	)
	if err := row.Scan(&pid, &threadID, &boardSlug, &author, &text, &createdAt, &imageURL, &parentID, &tripcode, &ipAddress); err != nil {
		return nil, err
	}
	post := entity.NewPost(threadID, boardSlug, author, text, imageURL, parentID, tripcode, ipAddress)
	post.SetID(pid)
	post.SetCreatedAt(createdAt)
	return post, nil
}

func (r *PostgresPostRepository) Create(ctx context.Context, post *entity.Post) error {
	row := r.pool.QueryRow(ctx, `INSERT INTO posts (thread_id, board_slug, author, text, created_at, image_url, parent_id, tripcode, ip_address) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9) RETURNING id`,
		post.ThreadID(), post.BoardSlug(), post.Author(), post.Text(), post.CreatedAt(), post.ImageURL(), post.ParentID(), post.Tripcode(), post.IPAddress())
	var id int64
	if err := row.Scan(&id); err != nil {
		return err
	}
	post.SetID(id)
	return nil
}

func (r *PostgresPostRepository) Delete(ctx context.Context, id int64) error {
	_, err := r.pool.Exec(ctx, `DELETE FROM posts WHERE id=$1`, id)
	return err
}

func (r *PostgresPostRepository) GetRecent(ctx context.Context, threadID int64, limit int) ([]*entity.Post, error) {
	rows, err := r.pool.Query(ctx, `SELECT id, thread_id, board_slug, author, text, created_at, image_url, parent_id, tripcode, ip_address FROM posts WHERE thread_id=$1 ORDER BY created_at DESC LIMIT $2`, threadID, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var posts []*entity.Post
	for rows.Next() {
		var id int64
		var threadID int64
		var boardSlug, author, text string
		var createdAt time.Time
		var imageURL *string
		var parentID *int64
		var tripcode *string
		var ipAddress string
		if err := rows.Scan(&id, &threadID, &boardSlug, &author, &text, &createdAt, &imageURL, &parentID, &tripcode, &ipAddress); err != nil {
			return nil, err
		}
		post := entity.NewPost(threadID, boardSlug, author, text, imageURL, parentID, tripcode, ipAddress)
		post.SetID(id)
		post.SetCreatedAt(createdAt)
		posts = append(posts, post)
	}
	return posts, nil
}
