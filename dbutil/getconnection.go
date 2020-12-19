package dbutil

import (
	"gorm.io/gorm"
	"sync"
)

var (
	db   *gorm.DB
	once sync.Once
)

func Pool(typeconnection string) *gorm.DB {
	once.Do(func() {
		conn := New()
		switch typeconnection {
		case MYSQL:
			db = conn.MSSQL()
		case MSSQL:
			db = conn.MSSQL()
		case PSQL:
			db = conn.PSQL()
		}
	})
	return db
}
