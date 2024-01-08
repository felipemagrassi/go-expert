//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/felipemagrassi/go-expert/di/product"
	"github.com/google/wire"
)

var setRepositoryDefaultSet = wire.NewSet(
	product.NewProductRepository,
	wire.Bind(
		new(product.ProductRepositoryInterface),
		new(*product.ProductRepository),
	),
)

func NewUseCase(db *sql.DB) *product.ProductUseCase {
	wire.Build(
		setRepositoryDefaultSet,
		product.NewProductUseCase,
	)

	return &product.ProductUseCase{}
}
