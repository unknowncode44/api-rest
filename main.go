package main

import (
	"fmt"
	"net/http"

	"github.com/unknowncode44/api-rest/internal/db"
	"github.com/unknowncode44/api-rest/internal/models"
	"github.com/unknowncode44/api-rest/internal/routes"
)

func main() {
	// inicializamos la db
	db.InitDbService()

	// creamos las tablas si no existen (usamos nuestros modelos)
	db.DBE.AutoMigrate(models.User{})
	db.DBE.AutoMigrate(models.Role{})
	db.DBE.AutoMigrate(models.Vehicle{})

	// inicializamos el router
	r := routes.InitRouter()

	// definimos el puerto y el host
	port := 3044
	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("Servidor Escuchando en http://localhost%s\n", addr)

	// corremos el metodo para que el servidor empiece a escuchar
	err := http.ListenAndServe(addr, r)

	// si hay error tiramos por consola
	if err != nil {
		panic(err)
	}

}
