package middleware

import (
	"context"
	"net/http"

	jwt "github.com/movie/JWT"
)

func RequireAuth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := r.Cookie("authorization")
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)

			return
		}

		claims, err := jwt.ValidateToken(token.Value)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotAcceptable)

			return
		}

		ctx := context.WithValue(r.Context(), "userKey", claims["user"])
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
