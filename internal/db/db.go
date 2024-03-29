package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// TODO: Guardar la configuracion de conexion a la db en una variable de entorno
var dsn string = "host=149.50.128.86 user=user password=password dbname=golang port=35000"

// para poder usar la db conectada es necesario crearla fuera de la funcion InitDB
// usamos el * para indicar que la variable db es un puntero que indica el valor de memoria donde esta guardada la variable
// ? Info sobre punteros: https://www.youtube.com/watch?v=gvjON1S0drk&t=67s
var db *gorm.DB

func InitDbService() {
	// definimos una variable para asignar errores si los hubiera
	var err error

	// instanciamos la conexion usando gorm y su metodo Open y el driver de postgres
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}) // <-- pasamos el dsn (database url etc.)

	// si hay errores los imprimimos por consola y salimos
	if err != nil {
		fmt.Fprintf(os.Stderr, "No se puede conectar la DB: %v\n", err)
		os.Exit(1)
	}

	// si no los hay, mostramos el mensaje de exito
	if db != nil {
		fmt.Printf("Conexion a DB Exitosa! \n")
	}

}
