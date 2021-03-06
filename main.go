package main

import (
	"github.com/Leonardo-Antonio/api-movies/dbutil"
	"github.com/Leonardo-Antonio/api-movies/models"
	"github.com/Leonardo-Antonio/api-movies/router"
	"github.com/Leonardo-Antonio/api-movies/storage"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"log"
)

func main() {
	db := dbutil.Pool(dbutil.MYSQL)
	err := db.AutoMigrate(
		&models.Categories{},
		&models.Authors{},
		&models.Movies{},
	)
	if err != nil {
		log.Fatal(err)
	}

	movieStorage := storage.NewMovie(db)
	authorStorage := storage.NewAuthor(db)
	categoryStorage := storage.NewCategory(db)

	e := echo.New()
	e.Use(middleware.CORS())
	router.Movie(movieStorage, e)
	router.Author(authorStorage, e)
	router.Category(categoryStorage, e)
	e.Logger.Fatal(e.Start(":8080"))

}

//c := storage.NewCategory(db)
//m := storage.NewMovie(db)
//a := storage.NewAuthor(db)
//
//err = categoryStorage.Create(&models.Categories{
//Category: "Terror",
//})
//err = categoryStorage.Create(&models.Categories{
//Category: "Comedy",
//})
//err = categoryStorage.Create(&models.Categories{
//Category: "Acción",
//})
//
//err = movieStorage.Create(&models.Movies{
//Name: "Saw I",
//Stars: 5,
//CategoriesID: 1,
//})
//err = movieStorage.Create(&models.Movies{
//Name: "Saw II",
//Stars: 4,
//CategoriesID: 1,
//})
//err = movieStorage.Create(&models.Movies{
//Name: "Saw III",
//Stars: 5,
//CategoriesID: 1,
//})
//err = authorStorage.Create(models.Authors{
//Name: "Leonardo",
//LastName: "Nolasco",
//Country: "Perú",
//MoviesID: 1,
//})
//err = authorStorage.Create(models.Authors{
//Name: "Alexandra",
//LastName: "Navarro",
//Country: "Perú",
//MoviesID: 1,
//})
//err = authorStorage.Create(models.Authors{
//Name: "Antonio",
//LastName: "Leyva",
//Country: "Perú",
//MoviesID: 2,
//})
//err = authorStorage.Create(models.Authors{
//Name: "Jaqueline",
//LastName: "Navarro",
//Country: "Perú",
//MoviesID: 2,
//})
