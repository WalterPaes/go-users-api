package repositories

import (
	"testing"

	"github.com/WalterPaes/go-users-api/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestUserRepository(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.User{})

	user, _ := entity.NewUser("John Doe", "j@j.com", "123456")

	r := NewUserRepository(db)

	t.Run("Should create an user", func(t *testing.T) {
		err := r.Create(user)
		if err != nil {
			t.Errorf("User.Create() error = %v", err)
		}
	})

	t.Run("Should find all users without pagination", func(t *testing.T) {
		users, err := r.FindAll(0, 0)
		if err != nil {
			t.Errorf("User.FindAll() error = %v", err)
		}

		assert.NotNil(t, users)
		assert.Greater(t, len(users), 0)
		assert.Equal(t, users[0].ID.String(), user.ID.String())
	})

	t.Run("Should find all users with pagination", func(t *testing.T) {
		user2, _ := entity.NewUser("Teste", "email@email.com", "123456")
		err := r.Create(user2)
		if err != nil {
			t.Errorf("User.Create() error = %v", err)
		}

		users, err := r.FindAll(2, 1)
		if err != nil {
			t.Errorf("User.FindAll() error = %v", err)
		}

		assert.NotNil(t, users)
		assert.Greater(t, len(users), 0)
		assert.Equal(t, users[0].ID.String(), user2.ID.String())
	})

	t.Run("Should find an user by id", func(t *testing.T) {
		result, err := r.FindById(user.ID)
		if err != nil {
			t.Errorf("User.FindById() error = %v", err)
		}

		assert.Equal(t, user.ID, result.ID)
		assert.Equal(t, user.Name, result.Name)
		assert.Equal(t, user.Email, result.Email)
		assert.True(t, result.ValidatePassword("123456"))
	})

	newUser := &entity.User{
		ID:    user.ID,
		Name:  "Editado",
		Email: "email@editado.com",
	}

	t.Run("Should update an user by id", func(t *testing.T) {
		err := r.Update(newUser)
		if err != nil {
			t.Errorf("User.Update() error = %v", err)
		}
	})

	t.Run("Should find an updated user by id", func(t *testing.T) {
		result, err := r.FindById(newUser.ID)
		if err != nil {
			t.Errorf("User.FindById() error = %v", err)
		}

		assert.Equal(t, newUser.ID, result.ID)
		assert.Equal(t, newUser.Name, result.Name)
		assert.Equal(t, newUser.Email, result.Email)
	})

	t.Run("Should delete an user by id", func(t *testing.T) {
		err := r.Delete(newUser.ID)
		if err != nil {
			t.Errorf("User.Delete() error = %v", err)
		}
	})
}
