package models

import "time"

type User struct {
	ID           uint64    `json:"id,omitempty"`
	Username     string    `json:"name,omitempty"`
	Nick         string    `json:"nick,omitempty"`
	Email        string    `json:"email,omitempty"`
	UserPassword string    `json:"password,omitempty"`
	CreatedAt    time.Time `json:"createdAt,omitempty"`
}
