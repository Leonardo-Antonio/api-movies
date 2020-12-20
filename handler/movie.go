package handler

import (
	"errors"
	"github.com/Leonardo-Antonio/api-movies/helpers"
	"github.com/Leonardo-Antonio/api-movies/models"
	"github.com/Leonardo-Antonio/api-movies/storage"
	"github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

type Movie struct {
	storage storage.IMovie
}

func NewMovie(storage storage.IMovie) *Movie {
	return &Movie{storage}
}

func (m *Movie) Create(c echo.Context) error {
	var model models.Movies
	err := c.Bind(&model)
	if err != nil {
		response := helpers.ResponseJSON(helpers.ERROR, "the structure is invalid", true, nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	err = m.storage.Create(&model)
	if err != nil {
		if errors.Is(err, helpers.ErrStars) || model.Name == "" {
			response := helpers.ResponseJSON(helpers.ERROR, err.Error(), true, nil)
			return c.JSON(http.StatusBadRequest, response)
		}
		e, ok := err.(*mysql.MySQLError)
		if !ok {
			response := helpers.ResponseJSON(helpers.ERROR, err.Error(), true, nil)
			return c.JSON(http.StatusInternalServerError, response)
		}
		if e.Number == 1062 || e.Number == 1452 {
			response := helpers.ResponseJSON(helpers.ERROR, e.Error(), true, nil)
			return c.JSON(http.StatusBadRequest, response)
		} else {
			response := helpers.ResponseJSON(helpers.ERROR, e.Error(), true, nil)
			return c.JSON(http.StatusInternalServerError, response)
		}
	}

	response := helpers.ResponseJSON(helpers.MESSAGE, "successfully created", false, nil)
	return c.JSON(http.StatusCreated, response)
}

func (m *Movie) Update(c echo.Context) error {
	var model models.Movies
	err := c.Bind(&model)
	if err != nil {
		response := helpers.ResponseJSON(helpers.ERROR, "the structure is invalid", true, nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	err = m.storage.Update(model)
	if err != nil {
		if errors.Is(err, helpers.ErrIDInvalid) || model.Name == "" {
			response := helpers.ResponseJSON(helpers.ERROR, err.Error(), true, nil)
			return c.JSON(http.StatusBadRequest, response)
		}
		e, ok := err.(*mysql.MySQLError)
		if !ok {
			response := helpers.ResponseJSON(helpers.ERROR, err.Error(), true, nil)
			return c.JSON(http.StatusInternalServerError, response)
		}
		if e.Number == 1062 || e.Number == 1452 {
			response := helpers.ResponseJSON(helpers.ERROR, e.Error(), true, nil)
			return c.JSON(http.StatusBadRequest, response)
		} else {
			response := helpers.ResponseJSON(helpers.ERROR, e.Error(), true, nil)
			return c.JSON(http.StatusInternalServerError, response)
		}
	}

	response := helpers.ResponseJSON(helpers.MESSAGE, "successfully updated", false, nil)
	return c.JSON(http.StatusOK, response)
}

func (m *Movie) Delete(c echo.Context) error {
	ID, err := strconv.Atoi(c.Param("ID"))
	if err != nil {
		response := helpers.ResponseJSON(helpers.ERROR, "parameter must be number", true, nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	err = m.storage.Delete(ID)
	if err != nil {
		if errors.Is(err, helpers.ErrIDInvalid) {
			response := helpers.ResponseJSON(helpers.ERROR, err.Error(), true, nil)
			return c.JSON(http.StatusBadRequest, response)
		} else {
			response := helpers.ResponseJSON(helpers.ERROR, err.Error(), true, nil)
			return c.JSON(http.StatusInternalServerError, response)
		}
	}

	response := helpers.ResponseJSON(helpers.MESSAGE, "successfully deleted", false, nil)
	return c.JSON(http.StatusOK, response)
}

func (m *Movie) GetAll(c echo.Context) error {
	data, err := m.storage.GetAll()

	if err != nil {
		e, ok := err.(*mysql.MySQLError)
		if !ok {
			response := helpers.ResponseJSON(helpers.ERROR, err.Error(), true, nil)
			return c.JSON(http.StatusInternalServerError, response)
		}
		if e.Number == 1062 || e.Number == 1452 {
			response := helpers.ResponseJSON(helpers.ERROR, e.Error(), true, nil)
			return c.JSON(http.StatusBadRequest, response)
		} else {
			response := helpers.ResponseJSON(helpers.ERROR, e.Error(), true, nil)
			return c.JSON(http.StatusInternalServerError, response)
		}
	}

	response := helpers.ResponseJSON(helpers.MESSAGE, "ok", false, data)
	return c.JSON(http.StatusOK, response)
}

func (m *Movie) GetByCategories(c echo.Context) error {
	ID, err := strconv.Atoi(c.Param("ID"))
	if err != nil {
		response := helpers.ResponseJSON(helpers.ERROR, "parameter must be number", true, nil)
		return c.JSON(http.StatusInternalServerError, response)
	}
	data, err := m.storage.GetByCategories(ID)

	if err != nil {
		e, ok := err.(*mysql.MySQLError)
		if !ok {
			response := helpers.ResponseJSON(helpers.ERROR, err.Error(), true, nil)
			return c.JSON(http.StatusInternalServerError, response)
		}
		if e.Number == 1062 || e.Number == 1452 {
			response := helpers.ResponseJSON(helpers.ERROR, e.Error(), true, nil)
			return c.JSON(http.StatusBadRequest, response)
		} else {
			response := helpers.ResponseJSON(helpers.ERROR, e.Error(), true, nil)
			return c.JSON(http.StatusInternalServerError, response)
		}
	}

	response := helpers.ResponseJSON(helpers.MESSAGE, "ok", false, data)
	return c.JSON(http.StatusOK, response)
}

func (m *Movie) GetByStars(c echo.Context) error {
	stars, err := strconv.Atoi(c.Param("stars"))
	if err != nil {
		response := helpers.ResponseJSON(helpers.ERROR, "parameter must be number", true, nil)
		return c.JSON(http.StatusInternalServerError, response)
	}
	data, err := m.storage.GetByStars(stars)

	if err != nil {
		if errors.Is(err, helpers.ErrStars) {
			response := helpers.ResponseJSON(helpers.ERROR, err.Error(), true, nil)
			return c.JSON(http.StatusBadRequest, response)
		}
		e, ok := err.(*mysql.MySQLError)
		if !ok {
			response := helpers.ResponseJSON(helpers.ERROR, err.Error(), true, nil)
			return c.JSON(http.StatusInternalServerError, response)
		}
		if e.Number == 1062 || e.Number == 1452 {
			response := helpers.ResponseJSON(helpers.ERROR, e.Error(), true, nil)
			return c.JSON(http.StatusBadRequest, response)
		} else {
			response := helpers.ResponseJSON(helpers.ERROR, e.Error(), true, nil)
			return c.JSON(http.StatusInternalServerError, response)
		}
	}

	response := helpers.ResponseJSON(helpers.MESSAGE, "ok", false, data)
	return c.JSON(http.StatusOK, response)
}
