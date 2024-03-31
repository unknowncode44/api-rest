package models

import "gorm.io/gorm"

type Vehicle struct {
	gorm.Model
	ID          int
	Brand       string `gorm:"type:varchar(100)"`
	Modelo      string
	PlateNumber string `gorm:"type:varchar(10); uniqueIndex"`
	Type        VehicleType
	UserID      uint
}
type VehicleType int

const (
	VehicleTypePickUp VehicleType = iota
	VehicleTypeWagon
	VehicleTypeCamper
	VehicleTypeAuto
	VehicleTypeFurgon
)

func (v VehicleType) String() string {
	switch v {
	case VehicleTypePickUp:
		return "PickUp"
	case VehicleTypeWagon:
		return "Wagon"
	case VehicleTypeCamper:
		return "Camper"
	case VehicleTypeAuto:
		return "Auto"
	case VehicleTypeFurgon:
		return "Furgon"
	}
	return "Unknown"
}
