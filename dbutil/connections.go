package dbutil

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type connection struct{}

func New() *connection {
	return &connection{}
}

func (c *connection) MYSQL() *gorm.DB {
	dsn := "leo:chester@tcp(127.0.0.1:3306)/db_movies?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func (c *connection) MSSQL() *gorm.DB {
	dsn := "leo:chester@tcp(127.0.0.1:3306)/db_movies?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func (c *connection) PSQL() *gorm.DB {
	dsn := "leo:chester@tcp(127.0.0.1:3306)/db_movies?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}
