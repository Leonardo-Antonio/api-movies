package handler

import (
	"github.com/Leonardo-Antonio/api-movies/helpers"
	"github.com/Leonardo-Antonio/api-movies/storage"
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
