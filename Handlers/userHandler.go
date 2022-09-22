package handlers

import (
	"encoding/json"
	"net/http"

	database "github.com/movie/Database"
	jwt "github.com/movie/JWT"
	models "github.com/movie/Models"
	"golang.org/x/crypto/bcrypt"
)

type result struct {
	Result string
}

func Register(w http.ResponseWriter, r *http.Request) {
	form := models.Form{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&form)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(form.Password), 10)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	db := database.InitDB()
	userResult := db.Create(&models.User{Username: form.Username, Password: hash})

	if userResult.Error != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(userResult.Error.Error())

		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result{Result: "success"})
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	form := models.Form{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&form)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err.Error())

		return
	}

	var user models.User
	db := database.InitDB()

	db.First(&user, "username = ?", form.Username)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err.Error())

		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(form.Password))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err.Error())

		return
	}

	tokenString, err := jwt.CreateToken(user.ID)
	if err != nil {

		http.Error(w, err.Error(), http.StatusNotFound)

		return
	}

	http.SetCookie(w, &http.Cookie{Name: "authorization", Value: tokenString, MaxAge: 3600 * 30, HttpOnly: false})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("success")
	// http.Redirect(w, r, "http://127.0.0.1:8080", http.StatusMovedPermanently)
}
