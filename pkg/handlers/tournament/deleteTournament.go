package tournament

import (
	"sport-tournaments/pkg/handlers"
	"sport-tournaments/pkg/models/databaseModels"
)

func DeleteTournament(id uint) {
	var tournament databaseModels.Tournament
	db := handlers.Database()
	db.First(&tournament, id)

	db.Delete(&tournament)
}
