// internal/service/user_service.go
package service

import (
	"github.com/izet28/user_service/internal/model"
	"github.com/izet28/user_service/internal/repository"
)

type UserService interface {
	CreateUser(user *model.User) (*model.User, error)
	GetUserByID(id int) (*model.User, error)
	GetAllUsers() ([]model.User, error)
	UpdateUser(user *model.User) (*model.User, error)
	DeleteUser(id int) error
}

type userService struct {
	repo repository.UserRepository
}

// NewUserService creates a new UserService.
func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(user *model.User) (*model.User, error) {
	return s.repo.CreateUser(user)
}

func (s *userService) GetUserByID(id int) (*model.User, error) {
	return s.repo.GetUserByID(id)
}

func (s *userService) GetAllUsers() ([]model.User, error) {
	return s.repo.GetAllUsers()
}

func (s *userService) UpdateUser(user *model.User) (*model.User, error) {
	return s.repo.UpdateUser(user)
}

func (s *userService) DeleteUser(id int) error {
	return s.repo.DeleteUser(id)
}
