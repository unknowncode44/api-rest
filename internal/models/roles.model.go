package models

import (
	"errors"

	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	ID   int      `gorm:"autoGenerate"`
	Type UserRole `gorm:"not null; uniqueIndex"`
}

// Definimos el tipo personalizado "RolUsuario" como un alias de string
type UserRole string

// Constantes que representan los diferentes roles
const (
	RolUsuarioAdministrador UserRole = "Administrator"
	RolUsuarioUsuario       UserRole = "User"
	RolUsuarioDesarrollador UserRole = "Developer"
	RolUsuarioControlador   UserRole = "Controller"
	RolUsuarioConductor     UserRole = "Driver"
	RolUsuarioRevendedor    UserRole = "Reseller"
)

// Función para obtener una lista de todos los roles
func AllRoles() []UserRole {
	return []UserRole{
		RolUsuarioAdministrador,
		RolUsuarioUsuario,
		RolUsuarioDesarrollador,
		RolUsuarioControlador,
		RolUsuarioConductor,
		RolUsuarioRevendedor,
	}
}

// Función para verificar si un rol es válido
func IsValidRole(role UserRole) error {

	for _, r := range AllRoles() {
		if r == role {
			return nil
		}
	}
	return errors.New("no es rol valido")
}
