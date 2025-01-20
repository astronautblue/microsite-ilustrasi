package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username  string `json:"username"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
}
