package tax

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTax(t *testing.T) {
	result, err := CalculateTax(1000)
	assert.Nil(t, err)
	assert.Equal(t, 10.0, result)

	result, err = CalculateTax(500)
	assert.Nil(t, err)
	assert.Equal(t, 5.0, result)

	result, err = CalculateTax(-1)
	assert.Error(t, err, "amount must be greater than 0")
	assert.Contains(t, err.Error(), "greater than 0")
	assert.Equal(t, 0.0, result)
}

func TestCalculateTax2(t *testing.T) {
	result, err := CalculateTax(1000)
	assert.Nil(t, err)
	assert.Equal(t, 10.0, result)

	result, err = CalculateTax(500)
	assert.Nil(t, err)
	assert.Equal(t, 5.0, result)

	result, err = CalculateTax(-1)
	assert.Error(t, err, "amount must be greater than 0")
	assert.Contains(t, err.Error(), "greater than 0")
	assert.Equal(t, 0.0, result)
}
