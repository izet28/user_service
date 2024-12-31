package services

import (
	"github.com/izet28/user_service/internal/models"
	"github.com/izet28/user_service/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	Repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (svc *UserService) GetAllUsers() ([]models.User, error) {
	return svc.Repo.GetAllUsers()
}

func (svc *UserService) CreateUser(user *models.User) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashedPassword)
	return svc.Repo.CreateUser(user)
}

func (svc *UserService) UpdateUser(id int, user *models.User) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashedPassword)
	return svc.Repo.UpdateUser(id, user)
}

func (svc *UserService) DeleteUser(id int) error {
	return svc.Repo.DeleteUser(id)
}
