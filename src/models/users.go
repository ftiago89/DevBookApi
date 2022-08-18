package models

import (
	"errors"
	"strings"
	"time"
)

type User struct {
	ID           uint64    `json:"id,omitempty"`
	Username     string    `json:"name,omitempty"`
	Nick         string    `json:"nick,omitempty"`
	Email        string    `json:"email,omitempty"`
	UserPassword string    `json:"password,omitempty"`
	CreatedAt    time.Time `json:"createdAt,omitempty"`
}

func (u *User) PrepareUser() error {
	u.format()

	if err := u.validate(); err != nil {
		return err
	}

	return nil
}

func (u *User) validate() error {
	if u.Username == "" {
		return errors.New("o nome é obrigatório e não pode ser vazio")
	}

	if u.Nick == "" {
		return errors.New("o nick é obrigatório e não pode ser vazio")
	}

	if u.Email == "" {
		return errors.New("o email é obrigatório e não pode ser vazio")
	}

	if u.UserPassword == "" {
		return errors.New("o password é obrigatório e não pode ser vazio")
	}

	return nil
}

func (u *User) format() {
	u.Username = strings.TrimSpace(u.Username)
	u.Nick = strings.TrimSpace(u.Nick)
	u.Email = strings.TrimSpace(u.Email)
}
