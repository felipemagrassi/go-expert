package product

import "database/sql"

type ProductRepository struct {
	db *sql.DB
}

type ProductRepositoryInterface interface {
	GetProduct(id string) (Product, error)
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db}
}

func (r *ProductRepository) GetProduct(id string) (Product, error) {
	return Product{
		ID:   id,
		Name: "Name",
	}, nil
}
