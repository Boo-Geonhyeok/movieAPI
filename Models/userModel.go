package models

import "gorm.io/gorm"

type Form struct {
	Username string
	Password string
}

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password []byte
}

type JWTtoken struct {
	Token string
}
