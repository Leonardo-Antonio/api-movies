package router

import (
	"github.com/Leonardo-Antonio/api-movies/handler"
	"github.com/Leonardo-Antonio/api-movies/storage"
	"github.com/labstack/echo"
)

func Category(storage storage.ICategory, e *echo.Echo) {
	hand := handler.NewCategory(storage)
	group := e.Group("/api/v1/categories")
	group.GET("", hand.GetAll)
	group.POST("", hand.Create)
	group.PUT("", hand.Update)
	group.DELETE("/:ID", hand.Delete)
}
