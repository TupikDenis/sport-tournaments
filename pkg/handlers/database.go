package handlers

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Database() *gorm.DB {
	dsn := "host=localhost user=postgres password=t7208588 dbname=sport_tournaments port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
