package user

import (
	"sport-tournaments/pkg/handlers"
	"sport-tournaments/pkg/models/databaseModels"
)

func UpdateUserPassword(id uint, newPassword string) {
	var user databaseModels.User

	db := handlers.Database()
	db.First(&user, id)

	user.Password = newPassword

	db.Save(&user)
}
