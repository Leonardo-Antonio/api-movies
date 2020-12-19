package models

import "gorm.io/gorm"

type Authors struct {
	gorm.Model
	Name     string `gorm:"type:varchar(50);not null" json:"name"`
	LastName string `gorm:"type:varchar(50);not null" json:"last_name"`
	Country  string `gorm:"type:varchar(20);not null" json:"country"`
	MoviesID uint   `json:"movies_id"`
}
