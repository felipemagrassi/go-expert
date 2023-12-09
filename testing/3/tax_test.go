package tax

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCalculateTax(t *testing.T) {
	t.Run("should return 5% when amount is less than 1000", func(t *testing.T) {
		amount := 500.0
		tax := CalculateTax(amount)
		assert.Equal(t, 5.0, tax)
	})
	t.Run("should return 10% when amount is greater than 1000", func(t *testing.T) {
		amount := 1500.0
		tax := CalculateTax(amount)
		assert.Equal(t, 10.0, tax)
	})
}

func TestCalculateTaxAndSave(t *testing.T) {
	repository := &TaxRepositoryMock{}
	repository.On("SaveTax", 5.0).Return(nil).Once()
	repository.On("SaveTax", 10.0).Return(errors.New("error saving tax")).Once()

	t.Run("should calculate tax and save", func(t *testing.T) {
		amount := 500.0
		error := CalculateTaxAndSave(amount, repository)

		assert.Equal(t, nil, error)
	})

	t.Run("should calculate tax and save", func(t *testing.T) {
		amount := 1500.0
		error := CalculateTaxAndSave(amount, repository)

		assert.Equal(t, errors.New("error saving tax"), error)
	})
}

type TaxRepositoryMock struct {
	mock.Mock
}

func (m *TaxRepositoryMock) SaveTax(amount float64) error {
	args := m.Called(amount)
	return args.Error(0)
}
