package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/unknowncode44/api-rest/internal/db"
	"github.com/unknowncode44/api-rest/internal/models"
)

func GetAllRoles(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get all roles"))
}

func PostOneRole(w http.ResponseWriter, r *http.Request) {
	var role models.Role
	json.NewDecoder(r.Body).Decode(&role)

	err := models.IsValidRole(role.Type)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Rol No Valido!"))
		return
	}

	createdUser := db.DBE.Create(&role)

	err = createdUser.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	} else {
		json.NewEncoder(w).Encode(&role)
	}

}
