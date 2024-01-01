package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/felipemagrassi/go-expert/graphql/graph"
	"github.com/felipemagrassi/go-expert/graphql/internal/database"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatalf("can't open database: %v", err)
	}

	defer db.Close()

	if port == "" {
		port = defaultPort
	}

	categoryDb := database.NewCategory(db)
	courseDb := database.NewCourse(db)
	resolver := &graph.Resolver{
		CategoryDB: categoryDb,
		CourseDB:   courseDb,
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
