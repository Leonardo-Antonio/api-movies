package storage

import (
	"github.com/Leonardo-Antonio/api-movies/helpers"
	"github.com/Leonardo-Antonio/api-movies/models"
	"gorm.io/gorm"
)

type (
	Category struct {
		db *gorm.DB
	}
	ICategory interface {
		Create(category *models.Categories) error
		GetAll() (categories []models.Categories, err error)
		Update(categories models.Categories) error
		Delete(ID int) error
	}
)

func NewCategory(db *gorm.DB) *Category {
	return &Category{db}
}

func (c *Category) Create(category *models.Categories) error {
	if category.Category == "" {
		return helpers.ErrNull("category")
	}
	err := c.db.Create(category).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *Category) GetAll() (categories []models.Categories, err error) {
	err = c.db.Select("*").Find(&categories).Error
	if err != nil {
		return
	}
	return
}

func (c *Category) Update(categories models.Categories) error {
	if categories.ID == 0 {
		return helpers.ErrIDInvalid
	}
	if categories.Category == "" {
		return helpers.ErrNull("categories")
	}
	err := c.db.Model(&models.Categories{}).Where("id = ?", categories.ID).
		Updates(&models.Categories{
			Category: categories.Category,
		}).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *Category) Delete(ID int) error {
	if ID == 0 {
		return helpers.ErrIDInvalid
	}
	err := c.db.Delete(&models.Categories{}, ID).Error
	if err != nil {
		return err
	}

	return nil
}
