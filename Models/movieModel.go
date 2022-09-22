package models

import (
	"gorm.io/gorm"
)

type WatchedList struct {
	gorm.Model
	Username string `json:"username"`
	MovieID  string `json:"id"`
}

type WishList struct {
	gorm.Model
	Username string `json:"username"`
	MovieID  string `json:"id"`
}

type Movie struct {
	ID string
}
