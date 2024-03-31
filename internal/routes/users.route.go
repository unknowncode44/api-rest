package routes

import "net/http"

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get users"))
}

func GetOneUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get user"))
}

func PostOneUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("post user"))
}

func PostBulkUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("post users"))
}

func DeleteOneUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("delete user"))
}

func DeleteBulkUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("delete users"))
}
