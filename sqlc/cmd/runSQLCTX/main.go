package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/felipemagrassi/go-expert/sqlc/internal/db"
	"github.com/google/uuid"

	_ "github.com/go-sql-driver/mysql"
)

type CourseDB struct {
	dbConnection *sql.DB
	*db.Queries
}

type CourseParams struct {
	ID          string
	Name        string
	Description sql.NullString
	Price       float64
}

type CategoryParams struct {
	ID          string
	Name        string
	Description sql.NullString
}

func (c *CourseDB) callTx(ctx context.Context, fn func(*db.Queries) error) error {
	tx, err := c.dbConnection.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := db.New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx error: %v, rb error: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

func NewCourseDB(dbConnection *sql.DB) *CourseDB {
	return &CourseDB{
		dbConnection: dbConnection,
		Queries:      db.New(dbConnection),
	}
}

func (c *CourseDB) CreateCourseAndCategory(ctx context.Context, argsCategory CategoryParams, argsCourse CourseParams) (db.Course, error) {
	err := c.callTx(ctx, func(q *db.Queries) error {
		err := q.CreateCategory(ctx, db.CreateCategoryParams{
			ID:          argsCategory.ID,
			Name:        argsCategory.Name,
			Description: argsCategory.Description,
		})

		if err != nil {
			return err
		}

		err = q.CreateCourse(ctx, db.CreateCourseParams{
			ID:          argsCourse.ID,
			Name:        argsCourse.Name,
			Description: argsCourse.Description,
			CategoryID:  argsCategory.ID,
			Price:       argsCourse.Price,
		})

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return db.Course{}, err
	}

	return db.Course{}, nil
}

func main() {
	ctx := context.Background()
	dbConnection, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	if err != nil {
		panic(err)
	}

	defer dbConnection.Close()

	queries := db.New(dbConnection)

	courseParams := CourseParams{
		ID:          uuid.New().String(),
		Name:        "Go",
		Description: sql.NullString{String: "Go!", Valid: true},
		Price:       100.50,
	}

	categoryParams := CategoryParams{
		ID:          uuid.New().String(),
		Name:        "Go Category",
		Description: sql.NullString{String: "Go Category!", Valid: true},
	}

	courseDB := NewCourseDB(dbConnection)
	_, err = courseDB.CreateCourseAndCategory(ctx, categoryParams, courseParams)
	if err != nil {
		panic(err)
	}

	courses, err := queries.ListCourses(ctx)
	if err != nil {
		panic(err)
	}

	for _, course := range courses {
		fmt.Printf("Category: %s | Course: %s %f %s\n", course.CategoryName, course.Name, course.Price, course.Description.String)
	}

}
