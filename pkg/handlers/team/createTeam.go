package team

import (
	"gorm.io/gorm"
	"sport-tournaments/pkg/handlers"
	"sport-tournaments/pkg/models/databaseModels"
)

func CreateTeam(teamName string, id uint) {
	db := handlers.Database()

	team := databaseModels.Team{
		Model:        gorm.Model{},
		TournamentId: id,
		TeamName:     teamName,
	}

	err := db.AutoMigrate(&databaseModels.Team{})
	if err != nil {
		return
	}

	db.Save(&team)
}
