package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("test", "test@test.com", "password")

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "test", user.Name)
	assert.Equal(t, "test@test.com", user.Email)
}

func TestComparePassword(t *testing.T) {
	user, err := NewUser("test", "test@test.com", "password123")
	assert.Nil(t, err)
	assert.NotNil(t, user)

	err = user.ComparePassword("password123")
	assert.Nil(t, err)

	err = user.ComparePassword("password")
	assert.NotNil(t, err)

}
