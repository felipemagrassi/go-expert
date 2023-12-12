package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	product, err := NewProduct("test", 100)

	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.NotEmpty(t, product.ID)
	assert.Equal(t, "test", product.Name)
	assert.Equal(t, 100, product.Price)
}

func TestNewProductInvalidPrice(t *testing.T) {
	product, err := NewProduct("test", -100)
	assert.NotNil(t, err)
	assert.Nil(t, product)
	assert.Error(t, err, ErrInvalidPrice)
}

func TestNewProductPriceIsRequired(t *testing.T) {
	product, err := NewProduct("test", 0)
	assert.NotNil(t, err)
	assert.Nil(t, product)
	assert.Error(t, err, ErrPriceIsRequired)
}

func TestNewProductNameIsRequired(t *testing.T) {
	product, err := NewProduct("", 100)
	assert.NotNil(t, err)
	assert.Nil(t, product)
	assert.Error(t, err, ErrNameIsRequired)
}
