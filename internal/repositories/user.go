package repositories

import (
	"github.com/WalterPaes/go-users-api/internal/entity"
	entityId "github.com/WalterPaes/go-users-api/pkg/entity"
	customError "github.com/WalterPaes/go-users-api/pkg/errors"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll(page, limit int) ([]entity.User, error)
	FindById(id entityId.ID) (*entity.User, error)
	Create(user *entity.User) error
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

func (r *User) FindAll(page, limit int) ([]entity.User, error) {
	var users []entity.User
	err := r.dbConn.Limit(limit).Offset((page - 1) * limit).Find(&users).Error
	return users, err
}

func (r *User) FindById(id entityId.ID) (*entity.User, error) {
	var user *entity.User
	result := r.dbConn.First(&user, "id = ?", id.String())
	if result.Error != nil {
		return nil, customError.New("UserRepositoryError::FindById", result.Error)
	}
	return user, nil
}

func (r *User) Create(user *entity.User) error {
	result := r.dbConn.Create(user)
	if result.Error != nil {
		return customError.New("UserRepositoryError::Create", result.Error)
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
		return customError.New("UserRepositoryError::Update", result.Error)
	}
	return nil
}

func (r *User) Delete(id entityId.ID) error {
	user, err := r.FindById(id)
	if err != nil {
		return err
	}

	result := r.dbConn.Delete(user)
	if result.Error != nil {
		return customError.New("UserRepositoryError::Delete", result.Error)
	}
	return nil
}
