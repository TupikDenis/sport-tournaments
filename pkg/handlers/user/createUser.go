package user

import (
	"gorm.io/gorm"
	"sport-tournaments/pkg/handlers"
	"sport-tournaments/pkg/models/databaseModels"
)

func CreateUser(username string, password string) {
	user := databaseModels.User{
		Model:    gorm.Model{},
		Username: username,
		Password: password,
		Role:     "user",
	}

	db := handlers.Database()
	err := db.AutoMigrate(&databaseModels.User{})
	if err != nil {
		return
	}
	
	db.Save(&user)
}
