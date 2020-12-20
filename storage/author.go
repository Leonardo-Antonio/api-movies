package storage

import (
	"github.com/Leonardo-Antonio/api-movies/helpers"
	"github.com/Leonardo-Antonio/api-movies/models"
	"gorm.io/gorm"
)

type (
	Author struct {
		db *gorm.DB
	}
	IAuthor interface {
		Create(actor models.Authors) (err error)
		Update(actor models.Authors) (err error)
		Delete(ID uint) (err error)
		GetByMovie(ID int) (movies models.Movies, err error)
	}
)

func NewAuthor(db *gorm.DB) *Author {
	return &Author{db}
}

func (a *Author) Create(actor models.Authors) (err error) {
	if actor.Name == "" {
		return helpers.ErrNull("name")
	}
	if actor.LastName == "" {
		return helpers.ErrNull("lastname")
	}
	if actor.Country == "" {
		return helpers.ErrNull("country")
	}

	err = a.db.Create(&actor).Error
	if err != nil {
		return
	}
	return
}

func (a *Author) Update(actor models.Authors) (err error) {
	if actor.Name == "" {
		return helpers.ErrNull("name")
	}
	if actor.LastName == "" {
		return helpers.ErrNull("lastname")
	}
	if actor.Country == "" {
		return helpers.ErrNull("country")
	}

	err = a.db.Model(&models.Authors{}).Where("id = ?", actor.ID).Updates(models.Authors{
		Name:     actor.Name,
		LastName: actor.LastName,
		Country:  actor.Country,
		MoviesID: actor.MoviesID,
	}).Error
	if err != nil {
		return
	}
	return
}

func (a *Author) Delete(ID uint) (err error) {
	if ID == 0 {
		return helpers.ErrIDInvalid
	}
	err = a.db.Delete(&models.Authors{}, ID).Error
	if err != nil {
		return
	}
	return nil
}

func (a *Author) GetByMovie(ID int) (movie models.Movies, err error) {
	if ID == 0 {
		return movie, helpers.ErrIDInvalid
	}
	rows, err := a.db.Model(&models.Movies{}).
		Select("authors.id, authors.created_at, "+
			"authors.updated_at, authors.deleted_at, "+
			"authors.name, authors.last_name, "+
			"authors.country, authors.movies_id, "+
			"movies.id, movies.created_at, "+
			"movies.updated_at, movies.deleted_at, "+
			"movies.name, movies.stars, movies.state, movies.categories_id").
		Joins("left join authors on movies.id = authors.movies_id").
		Where("movies.id = ?", ID).Rows()
	defer rows.Close()
	if err != nil {
		return
	}

	for rows.Next() {
		var authors models.Authors
		err := rows.Scan(
			&authors.ID, &authors.CreatedAt, &authors.UpdatedAt, &authors.DeletedAt,
			&authors.Name, &authors.LastName, &authors.Country, &authors.MoviesID,
			&movie.ID, &movie.CreatedAt, &movie.UpdatedAt, &movie.DeletedAt,
			&movie.Name, &movie.Stars, &movie.State, &movie.CategoriesID,
		)
		if err != nil {
			return models.Movies{}, err
		}
		movie.Actors = append(movie.Actors, authors)
	}
	if movie.ID == 0 {
		return movie, helpers.ErrMovieNotExist
	}
	return
}
