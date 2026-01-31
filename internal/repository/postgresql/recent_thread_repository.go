package postgresql

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vladimirfedunov/2chan-clone/internal/entity"
	"github.com/vladimirfedunov/2chan-clone/internal/repository"
)

type PostgresRecentThreadRepository struct {
	pool *pgxpool.Pool
}

func NewPostgresRecentThreadRepository(pool *pgxpool.Pool) repository.RecentThreadRepository {
	return &PostgresRecentThreadRepository{pool: pool}
}

func (r *PostgresRecentThreadRepository) GetRecent(ctx context.Context, limit int) ([]*entity.RecentThread, error) {
	// Исправленный запрос - берем только первый пост каждого треда
	query := `
        SELECT 
            t.id, 
            t.board_slug, 
            t.title, 
            COALESCE(first_post.text, '') as text,
            first_post.image_url, 
            t.created_at
        FROM threads t
        LEFT JOIN (
            SELECT DISTINCT ON (thread_id) 
                thread_id, text, image_url, created_at
            FROM posts 
            WHERE parent_id IS NULL
            ORDER BY thread_id, created_at ASC
        ) first_post ON t.id = first_post.thread_id
        ORDER BY t.created_at DESC
        LIMIT $1
    `

	log.Printf("Executing GetRecent query with limit: %d", limit)

	rows, err := r.pool.Query(ctx, query, limit)
	if err != nil {
		log.Printf("Query error: %v", err)
		return nil, err
	}
	defer rows.Close()

	var threads []*entity.RecentThread
	for rows.Next() {
		var (
			id        int64
			boardSlug string
			title     string
			text      string
			imageURL  *string
			createdAt time.Time
		)

		err := rows.Scan(&id, &boardSlug, &title, &text, &imageURL, &createdAt)
		if err != nil {
			log.Printf("Scan error: %v", err)
			continue
		}

		log.Printf("Found thread: ID=%d, Board=%s, Title=%s", id, boardSlug, title)

		thread := entity.NewRecentThread(id, boardSlug, title, text, imageURL, createdAt)
		threads = append(threads, thread)
	}

	log.Printf("Total threads found: %d", len(threads))
	return threads, nil
}

func (r *PostgresRecentThreadRepository) GetRecentByBoard(ctx context.Context, boardSlug string, limit int) ([]*entity.RecentThread, error) {
	query := `
        SELECT 
            t.id, 
            t.board_slug, 
            t.title, 
            COALESCE(first_post.text, '') as text,
            first_post.image_url, 
            t.created_at
        FROM threads t
        LEFT JOIN (
            SELECT DISTINCT ON (thread_id) 
                thread_id, text, image_url, created_at
            FROM posts 
            WHERE parent_id IS NULL
            ORDER BY thread_id, created_at ASC
        ) first_post ON t.id = first_post.thread_id
        WHERE t.board_slug = $1
        ORDER BY t.created_at DESC
        LIMIT $2
    `

	rows, err := r.pool.Query(ctx, query, boardSlug, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var threads []*entity.RecentThread
	for rows.Next() {
		var (
			id        int64
			boardSlug string
			title     string
			text      string
			imageURL  *string
			createdAt time.Time
		)

		err := rows.Scan(&id, &boardSlug, &title, &text, &imageURL, &createdAt)
		if err != nil {
			continue
		}

		thread := entity.NewRecentThread(id, boardSlug, title, text, imageURL, createdAt)
		threads = append(threads, thread)
	}

	return threads, nil
}
