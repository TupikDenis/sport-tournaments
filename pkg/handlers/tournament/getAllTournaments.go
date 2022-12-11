package tournament

import (
	"sport-tournaments/pkg/handlers"
	"sport-tournaments/pkg/models/databaseModels"
)

func GetAllTournaments() []databaseModels.TransformedTournament {
	var tournaments []databaseModels.Tournament
	var _tournaments []databaseModels.TransformedTournament

	db := handlers.Database()
	db.Find(&tournaments)

	for _, item := range tournaments {
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

	return _tournaments
}

func GetAllTournamentsByUserId(id uint) []databaseModels.TransformedTournament {
	var tournaments []databaseModels.Tournament
	var _tournaments []databaseModels.TransformedTournament

	db := handlers.Database()
	db.Find(&tournaments)

	for _, item := range tournaments {
		if item.UsernameId == id {
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

	return _tournaments
}
