package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	database "github.com/movie/Database"
	models "github.com/movie/Models"
	"gorm.io/gorm"
)

func AddWatchedMovie(w http.ResponseWriter, r *http.Request) {
	var watchedList models.WatchedList
	var user models.User
	movie := models.Movie{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&movie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)

		return
	}
	db := database.InitDB()
	result := db.First(&watchedList, "ID = ?", movie.ID)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		userKey := r.Context().Value("userKey")
		db.First(&user, userKey)
		result := db.Create(&models.WatchedList{Username: user.Username, MovieID: movie.ID})
		if result.Error != nil {
			http.Error(w, result.Error.Error(), http.StatusNotFound)

			return
		}
	} else {
		http.Error(w, "already exist", http.StatusConflict)

		return
	}
}

func AddWishMovie(w http.ResponseWriter, r *http.Request) {
	var wishList models.WishList
	var user models.User
	movie := models.Movie{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&movie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)

		return
	}
	db := database.InitDB()
	result := db.First(&wishList, "ID = ?", movie.ID)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		userKey := r.Context().Value("userKey")
		db.First(&user, userKey)
		result := db.Create(&models.WishList{Username: user.Username, MovieID: movie.ID})
		if result.Error != nil {
			http.Error(w, result.Error.Error(), http.StatusNotFound)

			return
		}
	} else {
		http.Error(w, "already exist", http.StatusConflict)

		return
	}
}

func DeleteWatchedMovie(w http.ResponseWriter, r *http.Request) {
	var watchedList models.WatchedList
	var user models.User
	movie := models.Movie{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&movie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)

		return
	}
	db := database.InitDB()
	result := db.First(&watchedList, "ID = ?", movie.ID)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		http.Error(w, "it doens't exist", http.StatusConflict)

		return
	} else {
		userKey := r.Context().Value("userKey")
		db.First(&user, userKey)
		result := db.Where("username = ?", user.Username).Delete(&watchedList)
		if result.Error != nil {
			http.Error(w, result.Error.Error(), http.StatusNotFound)

			return
		}
	}
}

func DeleteWishMovie(w http.ResponseWriter, r *http.Request) {
	var wishList models.WishList
	var user models.User
	movie := models.Movie{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&movie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)

		return
	}
	db := database.InitDB()
	result := db.First(&wishList, "ID = ?", movie.ID)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		http.Error(w, "it doens't exist", http.StatusConflict)

		return
	} else {
		userKey := r.Context().Value("userKey")
		db.First(&user, userKey)
		result := db.Where("username = ?", user.Username).Delete(&wishList)
		if result.Error != nil {
			http.Error(w, result.Error.Error(), http.StatusNotFound)

			return
		}
	}
}

func GetWatchedMovieList(w http.ResponseWriter, r *http.Request) {
	var watchedList []models.WatchedList
	var user models.User
	db := database.InitDB()
	userKey := r.Context().Value("userKey")
	db.First(&user, userKey)
	result := db.Where("username = ?", user.Username).Find(&watchedList)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusNotFound)

		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(watchedList)
}

func GetWishMovieList(w http.ResponseWriter, r *http.Request) {
	var wishList []models.WishList
	var user models.User
	db := database.InitDB()
	userKey := r.Context().Value("userKey")
	db.First(&user, userKey)
	result := db.Where("username = ?", user.Username).Find(&wishList)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusNotFound)

		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(wishList)
}
