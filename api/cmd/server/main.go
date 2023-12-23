package main

import (
	"net/http"

	"github.com/felipemagrassi/go-expert/api/configs"
	"github.com/felipemagrassi/go-expert/api/internal/entity"
	"github.com/felipemagrassi/go-expert/api/internal/infra/database"
	"github.com/felipemagrassi/go-expert/api/internal/infra/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	_, err := configs.LoadConfig(".")

	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.Product{}, &entity.User{})

	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	http.HandleFunc("/products", productHandler.CreateProduct)
	http.ListenAndServe(":8081", nil)
}
