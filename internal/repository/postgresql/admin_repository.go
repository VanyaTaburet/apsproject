package postgresql

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vladimirfedunov/2chan-clone/internal/entity"
	"github.com/vladimirfedunov/2chan-clone/internal/repository"
)

type PostgresAdminRepository struct {
	pool *pgxpool.Pool
}

func NewPostgresAdminRepository(pool *pgxpool.Pool) repository.AdminRepository {
	return &PostgresAdminRepository{pool: pool}
}

func (r *PostgresAdminRepository) GetByUsername(ctx context.Context, username string) (*entity.Admin, error) {
	row := r.pool.QueryRow(ctx, `SELECT id, username, password_hash, created_at, last_login FROM admins WHERE username=$1`, username)
	var (
		id          int64
		uname, hash string
		createdAt   time.Time
		lastLogin   *time.Time
	)
	if err := row.Scan(&id, &uname, &hash, &createdAt, &lastLogin); err != nil {
		return nil, err
	}
	admin := entity.NewAdmin(uname, hash)
	admin.SetID(id)
	admin.SetCreatedAt(createdAt)
	admin.SetLastLogin(lastLogin)
	return admin, nil
}

func (r *PostgresAdminRepository) Create(ctx context.Context, admin *entity.Admin) error {
	row := r.pool.QueryRow(ctx, `INSERT INTO admins (username, password_hash, created_at) VALUES ($1, $2, $3) RETURNING id`,
		admin.Username(), admin.PasswordHash(), admin.CreatedAt())
	var id int64
	if err := row.Scan(&id); err != nil {
		return err
	}
	admin.SetID(id)
	return nil
}

func (r *PostgresAdminRepository) UpdateLastLogin(ctx context.Context, id int64) error {
	_, err := r.pool.Exec(ctx, `UPDATE admins SET last_login=now() WHERE id=$1`, id)
	return err
}

func (r *PostgresAdminRepository) GetByID(ctx context.Context, id int64) (*entity.Admin, error) {
	row := r.pool.QueryRow(ctx, `SELECT id, username, password_hash, created_at, last_login FROM admins WHERE id=$1`, id)
	var (
		pid         int64
		uname, hash string
		createdAt   time.Time
		lastLogin   *time.Time
	)
	if err := row.Scan(&pid, &uname, &hash, &createdAt, &lastLogin); err != nil {
		return nil, err
	}
	admin := entity.NewAdmin(uname, hash)
	admin.SetID(pid)
	admin.SetCreatedAt(createdAt)
	admin.SetLastLogin(lastLogin)
	return admin, nil
}
