package repository

import (
	"context"
	"database/sql"

	"github.com/felipemagrassi/go-expert/unit-of-work/internal/db"
	"github.com/felipemagrassi/go-expert/unit-of-work/internal/entity"
)

type CourseRepositoryInterface interface {
	Insert(ctx context.Context, course entity.Course) error
}

type CourseRepository struct {
	DB      *sql.DB
	Queries *db.Queries
}

func NewCourseRepository(dbCon *sql.DB) *CourseRepository {
	return &CourseRepository{DB: dbCon, Queries: db.New(dbCon)}
}

func (r *CourseRepository) Insert(ctx context.Context, course entity.Course) error {
	return r.Queries.CreateCourse(ctx, db.CreateCourseParams{
		Name:       course.Name,
		CategoryID: int32(course.CategoryID),
	})
}
