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
	statement, err := u.db.Prepare(
		"insert into users (username, nick, email, userPassword) values (?, ?, ?, ?)",
	)

	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(user.Username, user.Nick, user.Email, user.UserPassword)
	if err != nil {
		return 0, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastId), nil
}
