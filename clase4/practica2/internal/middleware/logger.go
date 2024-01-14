package middleware

import (
	"fmt"
	"net/http"
	"time"
)

func Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("logger")

		start := time.Now()

		next.ServeHTTP(w, r)

		end := time.Since(start)
		fmt.Printf("request duration: %d", end.Nanoseconds())
	})
}
