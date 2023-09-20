package db

import (
	"WebApp/pkg/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDb(app *config.AppConfig) *gorm.DB {
	db, err := gorm.Open(postgres.Open(app.Dsn), &gorm.Config{})
	if err != nil {
		panic(err)
		// log.Println(err)
	}
	return db
}
