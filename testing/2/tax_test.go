package tax

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTax(t *testing.T) {
	assert.Equal(t, CalculateTax(1000), 5.0)
	assert.Equal(t, CalculateTax(1001), 10.0)

	assert.Positive(t, CalculateTax(1000))
}
