package router

import (
	"github.com/Leonardo-Antonio/api-movies/handler"
	"github.com/Leonardo-Antonio/api-movies/storage"
	"github.com/labstack/echo"
)

func Movie(movie storage.IMovie, e *echo.Echo) {
	hand := handler.NewMovie(movie)
	group := e.Group("/api/v1/movies")
	group.POST("", hand.Create)
	group.PUT("", hand.Update)
	group.DELETE("/:ID", hand.Delete)
	group.GET("/all", hand.GetAll)
	group.GET("/by/categories/:ID", hand.GetByCategories)
	group.GET("/by/stars/:stars", hand.GetByStars)
}
