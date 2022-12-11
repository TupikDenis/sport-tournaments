package match

import (
	"sport-tournaments/pkg/handlers"
	"sport-tournaments/pkg/models"
)

func GetMatchesByIdTournament(id uint) []models.TransformedFullMatch {
	var matches []models.FullMatch
	var _matches []models.TransformedFullMatch

	db := handlers.Database()
	db.Find(&matches)

	for _, item := range matches {
		if item.TournamentId == id {
			_matches = append(
				_matches, models.TransformedFullMatch{
					Id:           item.ID,
					TournamentId: item.TournamentId,
					Round:        item.Round,
					Robin:        item.Robin,
					Number:       item.Number,
					HomeTeam:     item.HomeTeam,
					HomeScore:    item.HomeScore,
					AwayTeam:     item.AwayTeam,
					AwayScore:    item.AwayScore,
				})
		}
	}

	return _matches
}
