package repository

import (
	"database/sql"

	"github.com/izet28/user_service/internal/models"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (repo *UserRepository) GetAllUsers() ([]models.User, error) {
	rows, err := repo.DB.Query("SELECT id, username, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (repo *UserRepository) GetUserByID(id int) (*models.User, error) {
	var user models.User
	err := repo.DB.QueryRow("SELECT id, username, email FROM users WHERE id = $1", id).Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepository) CreateUser(user *models.User) (*models.User, error) {
	err := repo.DB.QueryRow(
		"INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id",
		user.Username, user.Email, user.Password,
	).Scan(&user.ID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) UpdateUser(id int, user *models.User) (*models.User, error) {
	_, err := repo.DB.Exec(
		"UPDATE users SET username = $1, email = $2, password = $3 WHERE id = $4",
		user.Username, user.Email, user.Password, id,
	)
	if err != nil {
		return nil, err
	}
	user.ID = id
	return user, nil
}

func (repo *UserRepository) DeleteUser(id int) error {
	_, err := repo.DB.Exec("DELETE FROM users WHERE id = $1", id)
	return err
}
