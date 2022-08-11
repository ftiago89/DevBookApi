package repositories

import (
	"database/sql"
	"devbookapi/src/models"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{db}
}

func (u userRepository) Insert(user models.User) (uint64, error) {
	return 0, nil
}
