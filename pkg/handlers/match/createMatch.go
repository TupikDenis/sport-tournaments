package match

import (
	"gorm.io/gorm"
	"sport-tournaments/pkg/handlers"
	"sport-tournaments/pkg/models"
)

func CreateSchedule(matches []models.Match, robin int) {
	db := handlers.Database()

	var tournamentMatches []models.FullMatch

	for i := 0; i < len(matches); i++ {
		for k := 0; k < len(matches[i].Pair); k++ {
			tournamentMatches = append(tournamentMatches, models.FullMatch{
				Model:        gorm.Model{},
				TournamentId: matches[i].TournamentId,
				Round:        matches[i].Round,
				Robin:        robin,
				Number:       matches[i].Pair[k].Number,
				HomeTeam:     matches[i].Pair[k].HomeTeam,
				HomeScore:    matches[i].Pair[k].HomeScore,
				AwayTeam:     matches[i].Pair[k].AwayTeam,
				AwayScore:    matches[i].Pair[k].AwayScore,
			})
		}
	}

	err := db.AutoMigrate(&models.FullMatch{})
	if err != nil {
		return
	}

	db.Save(&tournamentMatches)
}
