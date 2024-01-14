package middleware

import (
	"fmt"
	"net/http"
)

var token string = "cluster"

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("auth")

		authHeader := r.Header.Get("Authorization")
		if authHeader != token {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
