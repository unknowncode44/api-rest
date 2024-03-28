package routes

import (
	"fmt"
	"net/http"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/api/data", apiDataHandler)

	return mux
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Bienvenidos al Home Page")
}

func apiDataHandler(w http.ResponseWriter, r *http.Request) {
	data := "datos"
	fmt.Fprintln(w, data)
}
