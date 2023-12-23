package database

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/felipemagrassi/go-expert/api/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file:test.db?mode=memory"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})

	product, _ := entity.NewProduct("TV", 1000)
	productDb := NewProduct(db)

	err = productDb.Create(product)
	assert.NoError(t, err)

	foundProduct, err := productDb.FindById(product.ID.String())
	assert.NoError(t, err)

	assert.Equal(t, foundProduct.ID, product.ID)
	assert.Equal(t, foundProduct.Name, product.Name)
	assert.Equal(t, foundProduct.Price, product.Price)
}

func TestUpdateProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file:test.db?mode=memory"), &gorm.Config{})

	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})

	product, err := entity.NewProduct("TV", 1000.0)
	assert.NoError(t, err)

	productDB := NewProduct(db)
	err = productDB.Create(product)
	assert.NoError(t, err)

	foundProduct, err := productDB.FindById(product.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, foundProduct.Price, product.Price)

	product.Price = 2000.0
	err = productDB.Update(product)
	assert.NoError(t, err)

	foundProduct, err = productDB.FindById(product.ID.String())

	assert.NoError(t, err)
	assert.Equal(t, foundProduct.Price, 2000.0)

}

func TestUpdateProductNotFound(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file:test.db?mode=memory"), &gorm.Config{})

	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})

	product, _ := entity.NewProduct("TV", 1000)
	productDB := NewProduct(db)

	product.Price = 2000
	err = productDB.Update(product)

	assert.NotNil(t, err)
}

func TestDeleteProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file:test.db?mode=memory"), &gorm.Config{})

	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})

	product, _ := entity.NewProduct("TV", 1000)
	productDb := NewProduct(db)

	err = productDb.Create(product)

	assert.Nil(t, err)

	foundProduct, err := productDb.FindById(product.ID.String())
	assert.Nil(t, err)
	assert.Equal(t, foundProduct.ID, product.ID)
	assert.Equal(t, foundProduct.Name, product.Name)
	assert.Equal(t, foundProduct.Price, product.Price)

	err = productDb.Delete(product.ID.String())
	assert.Nil(t, err)

	foundProduct, err = productDb.FindById(product.ID.String())
	assert.Nil(t, foundProduct)
	assert.NotNil(t, err)
}

func TestFindAllProducts(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file:test.db?mode=memory"), &gorm.Config{})

	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})

	for i := 1; i < 26; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("Product %d", i), rand.Float64()*1000)
		assert.NoError(t, err)
		db.Create(product)
	}
	productDb := NewProduct(db)

	products, err := productDb.FindAll(1, 10, "asc")

	assert.Nil(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, products[0].Name, "Product 1")
	assert.Equal(t, products[9].Name, "Product 10")

	products, err = productDb.FindAll(2, 10, "asc")
	assert.Nil(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, products[0].Name, "Product 11")
	assert.Equal(t, products[9].Name, "Product 20")

	products, err = productDb.FindAll(3, 10, "asc")
	assert.Nil(t, err)
	assert.Len(t, products, 5)
	assert.Equal(t, products[0].Name, "Product 21")
	assert.Equal(t, products[4].Name, "Product 25")

}
