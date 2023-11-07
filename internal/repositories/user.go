package repositories

import (
	"fmt"

	"github.com/WalterPaes/go-users-api/internal/entity"
	entityId "github.com/WalterPaes/go-users-api/pkg/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *entity.User) error
	FindById(id entityId.ID) (*entity.User, error)
	Update(user *entity.User) error
	Delete(id entityId.ID) error
}

type User struct {
	dbConn *gorm.DB
}

func NewUserRepository(db *gorm.DB) *User {
	return &User{
		dbConn: db,
	}
}

func (r *User) Create(user *entity.User) error {
	result := r.dbConn.Create(user)
	if result.Error != nil {
		return fmt.Errorf("[User Repository Error] %s", result.Error.Error())
	}
	return nil
}

func (r *User) Update(user *entity.User) error {
	_, err := r.FindById(user.ID)
	if err != nil {
		return err
	}

	result := r.dbConn.Save(user)
	if result.Error != nil {
		return fmt.Errorf("[User Repository Error] %s", result.Error.Error())
	}
	return nil
}

func (r *User) FindById(id entityId.ID) (*entity.User, error) {
	var user *entity.User
	result := r.dbConn.First(&user, "id = ?", id.String())
	if result.Error != nil {
		return nil, fmt.Errorf("[User Repository Error] %s", result.Error.Error())
	}
	return user, nil
}

func (r *User) Delete(id entityId.ID) error {
	user, err := r.FindById(id)
	if err != nil {
		return err
	}

	result := r.dbConn.Delete(user)
	if result.Error != nil {
		return fmt.Errorf("[User Repository Error] %s", result.Error.Error())
	}
	return nil
}
