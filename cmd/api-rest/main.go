package main

import (
	"fmt"
	"net/http"

	"github.com/unknowncode44/api-rest/internal/routes"
)

func main() {
	router := routes.NewRouter()
	port := 3044
	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("Servidor Escuchando en http://localhost%s\n", addr)
	err := http.ListenAndServe(addr, router)
	if err != nil {
		panic(err)
	}

}
