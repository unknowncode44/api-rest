package routes

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/unknowncode44/api-rest/internal/middleware"
)

func InitRouter() http.Handler {
	// creamos un nuevo router
	r := mux.NewRouter()

	//? ## RUTAS ##

	// Ruta Index
	r.HandleFunc("/", indexHandler)

	r.HandleFunc("/api/data", apiDataHandler)

	// RUTAS API USER
	r.HandleFunc("/api/users", GetAllUsers).Methods("GET")
	r.HandleFunc("/api/user/{id}", GetOneUser).Methods("GET")
	r.HandleFunc("/api/users", PostOneUser).Methods("POST")
	r.HandleFunc("/api/user/{id}", DeleteOneUser).Methods("DELETE")

	// RUTAS ROLES
	r.HandleFunc("/api/roles", middleware.RoleMiddleware(GetAllRoles)).Methods("GET")
	r.HandleFunc("/api/role", middleware.RoleMiddleware(PostOneRole)).Methods("POST")

	return r
}

func apiDataHandler(w http.ResponseWriter, r *http.Request) {
	data := "datos"
	fmt.Fprintln(w, data)
}
