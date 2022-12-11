package team

import (
	"sport-tournaments/pkg/handlers"
	"sport-tournaments/pkg/models/databaseModels"
)

func GetAllTournamentTeams(id uint) []databaseModels.TransformedTeam {
	var teams []databaseModels.Team
	var _teams []databaseModels.TransformedTeam

	db := handlers.Database()
	db.Find(&teams)

	for _, item := range teams {
		if item.TournamentId == id {
			_teams = append(
				_teams, databaseModels.TransformedTeam{
					Id:           item.ID,
					TournamentId: item.TournamentId,
					TeamName:     item.TeamName,
					Win:          item.Win,
					Draw:         item.Draw,
					Lose:         item.Lose,
					Score:        item.Score,
					Against:      item.Against,
				})
		}
	}

	return _teams
}
