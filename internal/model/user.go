package model

import (
	"time"
)

type User struct {
	UserId int `json:"user_id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Address string `json:"address"`
	CreatedAt time.Time `json:"created_at"`
}