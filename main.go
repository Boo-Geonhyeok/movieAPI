package main

import (
	"net/http"

	database "github.com/movie/Database"
	handlers "github.com/movie/Handlers"
	models "github.com/movie/Models"
	"github.com/movie/middleware"
)

func main() {
	db := database.InitDB()
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.WatchedList{})
	db.AutoMigrate(&models.WishList{})

	http.HandleFunc("/register", middleware.AllowCors(handlers.Register))
	http.HandleFunc("/login", middleware.AllowCors(handlers.Login))

	http.HandleFunc("/api/search", middleware.AllowCors(handlers.GetSearchedMovieList))
	http.HandleFunc("/api/post/wish", middleware.MultipleMiddleware(handlers.AddWishMovie, middleware.RequireAuth, middleware.AllowCors))
	http.HandleFunc("/api/post/watched", middleware.MultipleMiddleware(handlers.AddWatchedMovie, middleware.RequireAuth, middleware.AllowCors))
	http.HandleFunc("/api/delete/wish", middleware.MultipleMiddleware(handlers.DeleteWishMovie, middleware.RequireAuth, middleware.AllowCors))
	http.HandleFunc("/api/delete/watched", middleware.MultipleMiddleware(handlers.DeleteWatchedMovie, middleware.RequireAuth, middleware.AllowCors))
	http.HandleFunc("/api/get/wish", middleware.MultipleMiddleware(handlers.GetWishMovieList, middleware.RequireAuth, middleware.AllowCors))
	http.HandleFunc("/api/get/watched", middleware.MultipleMiddleware(handlers.GetWatchedMovieList, middleware.RequireAuth, middleware.AllowCors))

	http.ListenAndServe(":3000", nil)
}
