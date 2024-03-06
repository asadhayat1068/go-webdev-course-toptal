package main

import (
	"fmt"
	"net/http"
)

// Custom Middleware
func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("I am logging from custom middleware")
		next.ServeHTTP(w, r)
	})
}
