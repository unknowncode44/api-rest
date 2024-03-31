package middleware

import (
	"log"
	"net/http"
)

func RoleMiddleware(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Print(r.Method, " <-- Called in Route: --> ", r.URL.Path)
	}
}
