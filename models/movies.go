package models

import "gorm.io/gorm"

type Movies struct {
	gorm.Model
	Name         string `gorm:"type:varchar(40);unique;not null" json:"name"`
	Stars        int    `gorm:"size:4;not null" json:"stars"`
	State        bool   `gorm:"default:true;not null" json:"state"`
	Actors       []Authors
	CategoriesID uint `json:"categories_id"`
}
