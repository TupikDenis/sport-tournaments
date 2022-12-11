package olympic

import (
	"fmt"
	"math"
	"math/rand"
	"sport-tournaments/pkg/models"
	"sport-tournaments/pkg/models/databaseModels"
	"strconv"
	"time"
)

func Olympic(modelTeams []databaseModels.TransformedTeam, isMixed bool, id uint) []models.Match {
	if (math.Log2(float64(len(modelTeams))) - math.Floor(math.Log2(float64(len(modelTeams))))) != 0 {
		fmt.Println("Error!")
		return nil
	}

	var matches []models.Match

	checkRound := true
	numberMatch := 1

	for checkRound {
		var firstHalf []databaseModels.TransformedTeam
		var secondHalf []databaseModels.TransformedTeam

		if isMixed {
			shuffleTeams(modelTeams)
		}

		firstHalf = modelTeams[:len(modelTeams)/2]
		firstHalfTeam := addTeamsInSlice(firstHalf)

		secondHalf = modelTeams[len(modelTeams)/2:]
		secondHalfTeam := addTeamsInSlice(secondHalf)

		var pairs []models.Pair
		rand.Seed(time.Now().UnixNano())

		for k := 0; k < len(firstHalf); k++ {
			min := 0
			max := 100
			value := rand.Intn(max-min+1) + min
			if value < 50 {
				pairs = append(pairs, models.Pair{
					Number:    numberMatch,
					HomeTeam:  firstHalfTeam[k],
					HomeScore: 0,
					AwayTeam:  secondHalfTeam[k],
					AwayScore: 0,
				})
			} else {
				pairs = append(pairs, models.Pair{
					Number:    numberMatch,
					HomeTeam:  secondHalfTeam[k],
					HomeScore: 0,
					AwayTeam:  firstHalfTeam[k],
					AwayScore: 0,
				})
			}

			numberMatch++
		}

		var round string

		if len(firstHalf)%2 == 0 {
			round = "1/" + strconv.Itoa(len(modelTeams)/2)
		} else {
			round = "Final"
			checkRound = false
		}

		matches = append(matches, models.Match{
			TournamentId: id,
			Round:        round,
			Pair:         pairs,
		})

		var nextRoundTeam []databaseModels.TransformedTeam

		for i := 0; i < len(pairs); i++ {
			nextRoundTeam = append(nextRoundTeam, databaseModels.TransformedTeam{
				Id:           uint(i + 1),
				TournamentId: id,
				TeamName:     "Победитель матча №" + strconv.Itoa(pairs[i].Number), //matches[0].Pair[i].HomeTeam,
			})
		}

		modelTeams = nextRoundTeam
	}

	return matches
}

func shuffleTeams(teams []databaseModels.TransformedTeam) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(teams), func(i, j int) {
		teams[i],
			teams[j] = teams[j],
			teams[i]
	})
}

func addTeamsInSlice(modelTeams []databaseModels.TransformedTeam) []string {
	var teams []string

	for i := 0; i < len(modelTeams); i++ {
		teams = append(teams, modelTeams[i].TeamName)
	}

	return teams
}
