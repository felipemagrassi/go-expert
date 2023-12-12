package database

import (
	"testing"

	"github.com/felipemagrassi/go-expert/api/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file:test.db?mode=memory"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.User{})

	user, _ := entity.NewUser("Felipe", "Felipe@email.com", "password")
	userDb := NewUser(db)

	err = userDb.Create(user)

	assert.Nil(t, err)

	userFound, err := userDb.FindByEmail(user.Email)
	assert.Nil(t, err)
	assert.Equal(t, userFound.Email, user.Email)
	assert.Equal(t, userFound.Name, user.Name)
	assert.NotNil(t, userFound.Password)
}
