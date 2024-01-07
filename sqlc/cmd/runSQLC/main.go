package main

import (
	"context"
	"database/sql"

	"github.com/felipemagrassi/go-expert/sqlc/internal/db"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	ctx := context.Background()
	dbConnection, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	if err != nil {
		panic(err)
	}

	defer dbConnection.Close()

	queries := db.New(dbConnection)

	// for i := 0; i < 10; i++ {
	// 	err = queries.CreateCategory(ctx, db.CreateCategoryParams{
	// 		ID:          uuid.New().String(),
	// 		Name:        fmt.Sprintf("SQLC %d", i),
	// 		Description: sql.NullString{String: "SQLC is a tool to generate type safe Go from SQL", Valid: true},
	// 	})
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }
	//
	// categories, err := queries.ListCategories(ctx)
	// if err != nil {
	// 	panic(err)
	// }
	//
	// for i := 0; i < len(categories); i++ {
	// 	err = queries.UpdateCategory(ctx, db.UpdateCategoryParams{
	// 		ID:   categories[i].ID,
	// 		Name: fmt.Sprintf("SQLC %d updated", i),
	// 	})
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }
	//
	// categories, err = queries.ListCategories(ctx)
	// if err != nil {
	// 	panic(err)
	// }
	//
	// for _, catefories := range categories {
	// 	fmt.Println(catefories.Name)
	// }
	//
	// for i := 0; i < len(categories); i++ {
	// 	err = queries.DeleteCategory(ctx, categories[i].ID)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }

}
