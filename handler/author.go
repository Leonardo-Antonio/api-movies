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

type Author struct {
	storage storage.IAuthor
}

func NewAuthor(storage storage.IAuthor) *Author {
	return &Author{storage}
}

func (a *Author) GetByMovie(c echo.Context) error {
	ID, err := strconv.Atoi(c.Param("ID"))
	if err != nil {
		response := helpers.ResponseJSON(helpers.ERROR, "parameter must be number", true, nil)
		return c.JSON(http.StatusInternalServerError, response)
	}
	data, err := a.storage.GetByMovie(ID)
	if err != nil {
		response := helpers.ResponseJSON(helpers.ERROR, err.Error(), true, nil)
		return c.JSON(http.StatusInternalServerError, response)
	}
	response := helpers.ResponseJSON(helpers.MESSAGE, "OK", false, data)
	return c.JSON(http.StatusOK, response)
}

func (a *Author) Create(c echo.Context) error {
	var model models.Authors
	err := c.Bind(&model)
	if err != nil {
		response := helpers.ResponseJSON(helpers.ERROR, "the structure is invalid", true, nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	err = a.storage.Create(model)
	if err != nil {
		m, ok := err.(*mysql.MySQLError)
		if !ok {
			response := helpers.ResponseJSON(helpers.ERROR, err.Error(), true, nil)
			return c.JSON(http.StatusInternalServerError, response)
		}
		if m.Number == 1062 || m.Number == 1452 ||
			model.Name == "" || model.LastName == "" || model.Country == "" {
			response := helpers.ResponseJSON(helpers.ERROR, err.Error(), true, nil)
			return c.JSON(http.StatusBadRequest, response)
		} else {
			response := helpers.ResponseJSON(helpers.ERROR, err.Error(), true, nil)
			return c.JSON(http.StatusInternalServerError, response)
		}
	}

	response := helpers.ResponseJSON(helpers.MESSAGE, "successfully created", false, nil)
	return c.JSON(http.StatusCreated, response)
}

func (a *Author) Delete(c echo.Context) error {
	ID, err := strconv.Atoi(c.Param("ID"))
	if err != nil {
		response := helpers.ResponseJSON(helpers.ERROR, "parameter must be number", true, nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	err = a.storage.Delete(uint(ID))
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

func (a *Author) Update(c echo.Context) error {
	var model models.Authors
	err := c.Bind(&model)
	if err != nil {
		response := helpers.ResponseJSON(helpers.ERROR, "the structure is invalid", true, nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	err = a.storage.Update(model)
	if err != nil {
		m, ok := err.(*mysql.MySQLError)
		if !ok {
			response := helpers.ResponseJSON(helpers.ERROR, err.Error(), true, nil)
			return c.JSON(http.StatusInternalServerError, response)
		}
		if m.Number == 1062 || m.Number == 1452 ||
			model.Name == "" || model.LastName == "" || model.Country == "" {
			response := helpers.ResponseJSON(helpers.ERROR, err.Error(), true, nil)
			return c.JSON(http.StatusBadRequest, response)
		} else {
			response := helpers.ResponseJSON(helpers.ERROR, err.Error(), true, nil)
			return c.JSON(http.StatusInternalServerError, response)
		}
	}
	response := helpers.ResponseJSON(helpers.MESSAGE, "successfully updated", false, nil)
	return c.JSON(http.StatusOK, response)
}
