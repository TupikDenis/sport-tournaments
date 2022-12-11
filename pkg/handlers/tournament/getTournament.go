package tournament

import (
	"sport-tournaments/pkg/handlers"
	"sport-tournaments/pkg/models/databaseModels"
)

func GetTournamentById(id uint) databaseModels.TransformedTournament {
	var tournaments []databaseModels.Tournament
	var _tournaments []databaseModels.TransformedTournament
	var tournament databaseModels.TransformedTournament

	db := handlers.Database()
	db.Find(&tournaments)

	for _, item := range tournaments {
		if item.ID == id {
			_tournaments = append(
				_tournaments, databaseModels.TransformedTournament{
					Id:          item.ID,
					UsernameId:  item.UsernameId,
					Name:        item.Name,
					Description: item.Description,
					Sport:       item.Sport,
					System:      item.System,
				})
		}
	}

	if len(_tournaments) == 1 {
		tournament = _tournaments[0]
	}

	return tournament
}

func GetLastTournamentId() uint {
	var user databaseModels.Tournament

	db := handlers.Database()
	db.Last(&user)

	return user.ID
}
