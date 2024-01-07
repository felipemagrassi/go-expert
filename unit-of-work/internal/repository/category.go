package repository

import (
	"context"
	"database/sql"

	"github.com/felipemagrassi/go-expert/unit-of-work/internal/db"
	"github.com/felipemagrassi/go-expert/unit-of-work/internal/entity"
)

type CategoryRepositoryInterface interface {
	Insert(ctx context.Context, category entity.Category) error
}

type CategoryRepository struct {
	DB      *sql.DB
	Queries *db.Queries
}

func NewCategoryRepository(dbCon *sql.DB) *CategoryRepository {
	return &CategoryRepository{DB: dbCon, Queries: db.New(dbCon)}
}

func (r *CategoryRepository) Insert(ctx context.Context, category entity.Category) error {
	return r.Queries.CreateCategory(ctx, db.CreateCategoryParams{
		Name: category.Name,
	})
}
