package tournament

import (
	"github.com/jinzhu/gorm"
	"sport-tournaments/pkg/handlers"
	"sport-tournaments/pkg/models/databaseModels"
)

func CreateTournament(name string, description string, sport string, usernameId uint, system string) {
	db := handlers.Database()

	team := databaseModels.Tournament{
		Model:       gorm.Model{},
		UsernameId:  usernameId,
		Name:        name,
		Description: description,
		Sport:       sport,
		System:      system,
	}

	err := db.AutoMigrate(&databaseModels.Tournament{})
	if err != nil {
		return
	}

	db.Save(&team)

}
