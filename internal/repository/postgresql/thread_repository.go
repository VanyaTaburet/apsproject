package postgresql

import (
	"context"

	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vladimirfedunov/2chan-clone/internal/entity"
	"github.com/vladimirfedunov/2chan-clone/internal/repository"
)

type PostgresThreadRepository struct {
	pool *pgxpool.Pool
}

func NewPostgresThreadRepository(pool *pgxpool.Pool) repository.ThreadRepository {
	return &PostgresThreadRepository{pool: pool}
}

func (r *PostgresThreadRepository) GetByBoard(ctx context.Context, boardSlug string) ([]*entity.Thread, error) {
	rows, err := r.pool.Query(ctx, `SELECT id, board_slug, title, author, created_at, sticky, locked, bump_at FROM threads WHERE board_slug=$1 ORDER BY bump_at DESC`, boardSlug)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var threads []*entity.Thread
	for rows.Next() {
		var id int64
		var boardSlug, title, author string
		var createdAt, bumpAt time.Time
		var sticky, locked bool
		if err := rows.Scan(&id, &boardSlug, &title, &author, &createdAt, &sticky, &locked, &bumpAt); err != nil {
			return nil, err
		}
		thread := entity.NewThread(boardSlug, title, author, sticky, locked)
		thread.SetID(id)
		thread.SetCreatedAt(createdAt)
		thread.SetBumpAt(bumpAt)
		threads = append(threads, thread)
	}
	return threads, nil
}

func (r *PostgresThreadRepository) GetByID(ctx context.Context, id int64) (*entity.Thread, error) {
	row := r.pool.QueryRow(ctx, `SELECT id, board_slug, title, author, created_at, sticky, locked, bump_at FROM threads WHERE id=$1`, id)
	var tid int64
	var boardSlug, title, author string
	var createdAt, bumpAt time.Time
	var sticky, locked bool
	if err := row.Scan(&tid, &boardSlug, &title, &author, &createdAt, &sticky, &locked, &bumpAt); err != nil {
		return nil, err
	}
	thread := entity.NewThread(boardSlug, title, author, sticky, locked)
	thread.SetID(tid)
	thread.SetCreatedAt(createdAt)
	thread.SetBumpAt(bumpAt)
	return thread, nil
}

func (r *PostgresThreadRepository) Create(ctx context.Context, thread *entity.Thread) error {
	row := r.pool.QueryRow(ctx, `INSERT INTO threads (board_slug, title, author, created_at, sticky, locked, bump_at) VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING id`,
		thread.BoardSlug(), thread.Title(), thread.Author(), thread.CreatedAt(), thread.Sticky(), thread.Locked(), thread.BumpAt())
	var id int64
	if err := row.Scan(&id); err != nil {
		return err
	}
	thread.SetID(id)
	return nil
}

func (r *PostgresThreadRepository) Update(ctx context.Context, thread *entity.Thread) error {
	_, err := r.pool.Exec(ctx, `UPDATE threads SET title=$1, author=$2, sticky=$3, locked=$4, bump_at=$5 WHERE id=$6`,
		thread.Title(), thread.Author(), thread.Sticky(), thread.Locked(), thread.BumpAt(), thread.ID())
	return err
}

func (r *PostgresThreadRepository) Delete(ctx context.Context, id int64) error {
	_, err := r.pool.Exec(ctx, `DELETE FROM threads WHERE id=$1`, id)
	return err
}

func (r *PostgresThreadRepository) GetRecent(ctx context.Context, boardSlug string, limit int) ([]*entity.Thread, error) {
	var rows pgx.Rows
	var err error
	if limit > 0 {
		rows, err = r.pool.Query(ctx, `SELECT id, board_slug, title, author, created_at, sticky, locked, bump_at FROM threads WHERE board_slug=$1 ORDER BY bump_at DESC LIMIT $2`, boardSlug, limit)
	} else {
		rows, err = r.pool.Query(ctx, `SELECT id, board_slug, title, author, created_at, sticky, locked, bump_at FROM threads WHERE board_slug=$1 ORDER BY bump_at DESC`, boardSlug)
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var threads []*entity.Thread
	for rows.Next() {
		var id int64
		var boardSlug, title, author string
		var createdAt, bumpAt time.Time
		var sticky, locked bool
		if err := rows.Scan(&id, &boardSlug, &title, &author, &createdAt, &sticky, &locked, &bumpAt); err != nil {
			return nil, err
		}
		thread := entity.NewThread(boardSlug, title, author, sticky, locked)
		thread.SetID(id)
		thread.SetCreatedAt(createdAt)
		thread.SetBumpAt(bumpAt)
		threads = append(threads, thread)
	}
	return threads, nil
}

func (r *PostgresThreadRepository) GetAllThreads(ctx context.Context) ([]*entity.Thread, error) {
	rows, err := r.pool.Query(ctx, `SELECT id, board_slug, title, author, created_at, sticky, locked FROM threads ORDER BY created_at DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var threads []*entity.Thread
	for rows.Next() {
		var id int64
		var boardSlug, title, author string
		var createdAt time.Time
		var sticky, locked bool
		if err := rows.Scan(&id, &boardSlug, &title, &author, &createdAt, &sticky, &locked); err != nil {
			return nil, err
		}
		thread := entity.NewThread(boardSlug, title, author, sticky, locked)
		thread.SetID(id)
		thread.SetCreatedAt(createdAt)
		thread.SetSticky(sticky)
		thread.SetLocked(locked)
		threads = append(threads, thread)
	}
	return threads, nil
}
