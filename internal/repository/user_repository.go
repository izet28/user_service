// internal/repository/user_repository.go
package repository

import (
	"github.com/izet28/user_service/internal/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // PostgreSQL dialect
)

type UserRepository interface {
	CreateUser(user *model.User) (*model.User, error)
	GetUserByID(id int) (*model.User, error)
	GetAllUsers() ([]model.User, error)
	UpdateUser(user *model.User) (*model.User, error)
	DeleteUser(id int) error
}

type UserRepositoryGORM struct {
	DB *gorm.DB
}

// NewUserRepository creates a new instance of UserRepository using GORM.
func NewUserRepository(db *gorm.DB) *UserRepositoryGORM {
	return &UserRepositoryGORM{DB: db}
}

// CreateUser creates a new user in the database.
func (r *UserRepositoryGORM) CreateUser(user *model.User) (*model.User, error) {
	if err := r.DB.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// GetUserByID retrieves a user by ID from the database.
func (r *UserRepositoryGORM) GetUserByID(id int) (*model.User, error) {
	var user model.User
	if err := r.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetAllUsers retrieves all users from the database.
func (r *UserRepositoryGORM) GetAllUsers() ([]model.User, error) {
	var users []model.User
	if err := r.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// UpdateUser updates an existing user.
func (r *UserRepositoryGORM) UpdateUser(user *model.User) (*model.User, error) {
	if err := r.DB.Save(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// DeleteUser deletes a user by ID.
func (r *UserRepositoryGORM) DeleteUser(id int) error {
	if err := r.DB.Delete(&model.User{}, id).Error; err != nil {
		return err
	}
	return nil
}
