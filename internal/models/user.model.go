package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	ID        uint   `gorm:"primaryKey; autoIncrement; "`
	Name      string `gorm:"not null; uniqueIndex"`
	Email     string `gorm:"not null; uniqueIndex"`
	Password  string `gorm:"not null"`
	RoleID    int
	Role      Role
	Active    bool `gorm:"default:false"`
	Vehicles  []Vehicle
	CreatedAt time.Time
	UpdatedAt time.Time
}
