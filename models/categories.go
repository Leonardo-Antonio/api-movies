package models

import "gorm.io/gorm"

type Categories struct {
	gorm.Model
	Category string `gorm:"type:varchar(30);unique;not null" json:"category"`
	Movies   Movies
}
