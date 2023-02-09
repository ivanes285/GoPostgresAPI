package models

import "gorm.io/gorm"

type User struct {
	gorm.Model //para que sepa  y tranforme en tabla de la base de datos

	FirstName string `gorm:"not null;unique"  json:"firstname"` //? json:"firstname" establezco este nombre para enviarlo desde el front
	LastName  string `gorm:"not null;"        json:"lastname"`
	Email     string `gorm:"not null;unique"  json:"email"`
	Tasks     []Task `json:"tasks"`
}
