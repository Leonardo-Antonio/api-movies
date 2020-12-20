package storage

import (
	"github.com/Leonardo-Antonio/api-movies/helpers"
	"github.com/Leonardo-Antonio/api-movies/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type (
	Movie struct {
		db *gorm.DB
	}
	IMovie interface {
		Create(movie *models.Movies) error
		Update(movie models.Movies) (err error)
		Delete(movie *models.Movies) (err error)
		GetAll() (movies []models.Movies, err error)
		GetByCategories(IDCategory int) (movies []models.Movies, err error)
		GetByStars(stars int) (movies []models.Movies, err error)
	}
)

func NewMovie(db *gorm.DB) *Movie {
	return &Movie{db}
}

func (m *Movie) Create(movie *models.Movies) error {
	if movie.Stars > 5 || movie.Stars < 0 {
		return helpers.ErrStars
	}
	if movie.Name == "" {
		return helpers.ErrNull("name")
	}
	err := m.db.Create(movie).Error
	if err != nil {
		return err
	}
	return nil
}

func (m *Movie) Update(movie models.Movies) (err error) {
	if movie.Name == "" {
		return helpers.ErrNull("name")
	}
	if movie.ID == 0 {
		return helpers.ErrIDInvalid
	}

	err = m.db.Model(&models.Movies{}).
		Where("id = ?", movie.ID).
		Updates(&models.Movies{
			Name:         movie.Name,
			Stars:        movie.Stars,
			State:        movie.State,
			CategoriesID: movie.CategoriesID,
		}).Error

	if err != nil {
		return
	}
	return
}

func (m *Movie) Delete(movie *models.Movies) (err error) {
	if movie.ID == 0 {
		return helpers.ErrIDInvalid
	}
	err = m.db.Delete(movie).Error
	if err != nil {
		return
	}
	return
}

func (m *Movie) GetAll() (movies []models.Movies, err error) {
	err = m.db.Preload(clause.Associations).Find(&movies).Error
	if err != nil {
		return
	}
	return
}

func (m *Movie) GetByCategories(ID int) (movies []models.Movies, err error) {
	err = m.db.Preload(clause.Associations).Find(&movies, "categories_id = ?", ID).Error
	if err != nil {
		return
	}
	return
}

func (m *Movie) GetByStars(stars int) (movies []models.Movies, err error) {
	if stars > 5 || stars < 0 {
		return movies, helpers.ErrStars
	}
	err = m.db.Preload(clause.Associations).Find(&movies, "stars = ?", stars).Error
	if err != nil {
		return
	}
	return
}
