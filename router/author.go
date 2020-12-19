package router

import (
	"github.com/Leonardo-Antonio/api-movies/handler"
	"github.com/Leonardo-Antonio/api-movies/storage"
	"github.com/labstack/echo"
)

func Author(storage storage.IAuthor, e *echo.Echo) {
	hand := handler.NewAuthor(storage)
	group := e.Group("/api/v1/authors")
	group.GET("/:ID", hand.GetByMovie)
}
