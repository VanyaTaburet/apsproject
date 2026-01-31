package postgresql

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vladimirfedunov/2chan-clone/internal/entity"
	"github.com/vladimirfedunov/2chan-clone/internal/repository"
)

type PostgresBoardRepository struct {
	pool *pgxpool.Pool
}

func NewPostgresBoardRepository(pool *pgxpool.Pool) repository.BoardRepository {
	return &PostgresBoardRepository{pool: pool}
}

func (r *PostgresBoardRepository) GetAll(ctx context.Context) ([]*entity.Board, error) {
	rows, err := r.pool.Query(ctx, `SELECT slug, name, description FROM boards ORDER BY slug`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var boards []*entity.Board
	for rows.Next() {
		var slug, name, description string
		if err := rows.Scan(&slug, &name, &description); err != nil {
			return nil, err
		}
		board := entity.NewBoard(slug, name, description)
		boards = append(boards, board)
	}
	return boards, nil
}

func (r *PostgresBoardRepository) GetBySlug(ctx context.Context, slug string) (*entity.Board, error) {
	row := r.pool.QueryRow(ctx, `SELECT slug, name, description FROM boards WHERE slug=$1`, slug)
	var s, name, description string
	if err := row.Scan(&s, &name, &description); err != nil {
		return nil, err
	}
	return entity.NewBoard(s, name, description), nil
}

func (r *PostgresBoardRepository) Create(ctx context.Context, board *entity.Board) error {
	_, err := r.pool.Exec(ctx, `INSERT INTO boards (slug, name, description) VALUES ($1,$2,$3)`,
		board.Slug(), board.Name(), board.Description())
	return err
}

func (r *PostgresBoardRepository) Update(ctx context.Context, board *entity.Board) error {
	_, err := r.pool.Exec(ctx, `UPDATE boards SET name=$1, description=$2 WHERE slug=$3`,
		board.Name(), board.Description(), board.Slug())
	return err
}

func (r *PostgresBoardRepository) Delete(ctx context.Context, slug string) error {
	_, err := r.pool.Exec(ctx, `DELETE FROM boards WHERE slug=$1`, slug)
	return err
}
