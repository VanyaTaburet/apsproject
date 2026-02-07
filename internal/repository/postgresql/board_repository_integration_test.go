//go:build integration

package postgresql

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/vladimirfedunov/2chan-clone/internal/entity"
)

func setupTestRepo(t *testing.T) (*PostgresBoardRepository, func()) {
	t.Helper()

	dsn := os.Getenv("TEST_DB_DSN")
	require.NotEmpty(t, dsn, "TEST_DB_DSN is not set")

	ctx := context.Background()

	pool, err := NewPostgresPool(ctx, dsn)
	require.NoError(t, err)

	repo := &PostgresBoardRepository{pool: pool}

	cleanup := func() {
		_, err := pool.Exec(ctx, `TRUNCATE TABLE boards`)
		require.NoError(t, err)
		pool.Close()
	}

	return repo, cleanup
}

func TestPostgresBoardRepository_CRUD(t *testing.T) {
	ctx := context.Background()

	repo, cleanup := setupTestRepo(t)
	defer cleanup()

	board := entity.NewBoard("tech", "Technology", "Tech board")

	// Create
	err := repo.Create(ctx, board)
	require.NoError(t, err)

	// GetBySlug
	got, err := repo.GetBySlug(ctx, "tech")
	require.NoError(t, err)
	assert.Equal(t, board.Slug(), got.Slug())
	assert.Equal(t, board.Name(), got.Name())
	assert.Equal(t, board.Description(), got.Description())

	// GetAll
	all, err := repo.GetAll(ctx)
	require.NoError(t, err)
	require.Len(t, all, 1)

	// Update
	updated := entity.NewBoard("tech", "Tech Updated", "Updated description")
	err = repo.Update(ctx, updated)
	require.NoError(t, err)

	got, err = repo.GetBySlug(ctx, "tech")
	require.NoError(t, err)
	assert.Equal(t, "Tech Updated", got.Name())
	assert.Equal(t, "Updated description", got.Description())

	// Delete
	err = repo.Delete(ctx, "tech")
	require.NoError(t, err)

	all, err = repo.GetAll(ctx)
	require.NoError(t, err)
	assert.Len(t, all, 0)
}
