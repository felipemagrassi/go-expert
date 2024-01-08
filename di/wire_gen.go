// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"database/sql"
	"github.com/felipemagrassi/go-expert/di/product"
	"github.com/google/wire"
)

import (
	_ "github.com/mattn/go-sqlite3"
)

// Injectors from wire.go:

func NewUseCase(db *sql.DB) *product.ProductUseCase {
	productRepository := product.NewProductRepository(db)
	productUseCase := product.NewProductUseCase(productRepository)
	return productUseCase
}

// wire.go:

var setRepositoryDefaultSet = wire.NewSet(product.NewProductRepository, wire.Bind(
	new(product.ProductRepositoryInterface),
	new(*product.ProductRepository),
),
)
