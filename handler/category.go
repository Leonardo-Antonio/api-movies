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

type Category struct {
	storage storage.ICategory
}

func NewCategory(storage storage.ICategory) *Category {
	return &Category{storage}
}

func (ct *Category) GetAll(c echo.Context) error {
	data, err := ct.storage.GetAll()
	if err != nil {
		response := helpers.ResponseJSON(helpers.ERROR, err.Error(), true, data)
		return c.JSON(http.StatusBadRequest, response)
	}
	response := helpers.ResponseJSON(helpers.MESSAGE, "OK", false, data)
	return c.JSON(http.StatusOK, response)
}

func (ct *Category) Create(c echo.Context) error {
	var model models.Categories
	err := c.Bind(&model)
	if err != nil {
		response := helpers.ResponseJSON(helpers.ERROR, "the structure is invalid", true, nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	err = ct.storage.Create(&model)
	if err != nil {
		m, ok := err.(*mysql.MySQLError)
		if !ok {
			response := helpers.ResponseJSON(helpers.ERROR, err.Error(), true, nil)
			return c.JSON(http.StatusInternalServerError, response)
		}
		if m.Number == 1062 || model.Category == "" {
			response := helpers.ResponseJSON(helpers.ERROR, err.Error(), true, nil)
			return c.JSON(http.StatusBadRequest, response)
		} else {
			response := helpers.ResponseJSON(helpers.ERROR, "the structure is invalid", true, nil)
			return c.JSON(http.StatusInternalServerError, response)
		}
	}

	response := helpers.ResponseJSON(helpers.MESSAGE, "successfully created", false, nil)
	return c.JSON(http.StatusCreated, response)
}

func (ct *Category) Update(c echo.Context) error {
	var model models.Categories
	err := c.Bind(&model)
	if err != nil {
		response := helpers.ResponseJSON(helpers.ERROR, "the structure is invalid", true, nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	err = ct.storage.Update(model)
	if err != nil {
		m, ok := err.(*mysql.MySQLError)
		if !ok {
			response := helpers.ResponseJSON(helpers.ERROR, err.Error(), true, nil)
			return c.JSON(http.StatusInternalServerError, response)
		}
		if m.Number == 1062 || model.Category == "" {
			response := helpers.ResponseJSON(helpers.ERROR, err.Error(), true, nil)
			return c.JSON(http.StatusBadRequest, response)
		} else {
			response := helpers.ResponseJSON(helpers.ERROR, "the structure is invalid", true, nil)
			return c.JSON(http.StatusInternalServerError, response)
		}
	}

	response := helpers.ResponseJSON(helpers.MESSAGE, "successfully updated", false, nil)
	return c.JSON(http.StatusOK, response)
}

func (ct *Category) Delete(c echo.Context) error {
	ID, err := strconv.Atoi(c.Param("ID"))
	if err != nil {
		response := helpers.ResponseJSON(helpers.ERROR, "parameter must be number", true, nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	err = ct.storage.Delete(ID)
	if err != nil {
		if errors.Is(err, helpers.ErrIDInvalid) {
			response := helpers.ResponseJSON(helpers.ERROR, err.Error(), true, nil)
			return c.JSON(http.StatusBadRequest, response)
		} else {
			response := helpers.ResponseJSON(helpers.ERROR, err.Error(), true, nil)
			return c.JSON(http.StatusInternalServerError, response)
		}
	}
	response := helpers.ResponseJSON(helpers.ERROR, "successfully deleted", false, nil)
	return c.JSON(http.StatusOK, response)
}
