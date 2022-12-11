package robin_round

import (
	"encoding/json"
	rand "math/rand"
	"sport-tournaments/pkg/models"
	"sport-tournaments/pkg/models/databaseModels"
	"strconv"
	"time"
)

func RobinRound(modelTeams []databaseModels.TransformedTeam, robin int, mixed bool, id uint) []models.Match {
	if mixed {
		shuffleTeams(modelTeams)
	}

	teams := convertTeamsInString(modelTeams)
	matches := createSchedule(teams, robin, id)
	//matchesJSON := madeJSON(matches)
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

func convertTeamsInString(modelTeams []databaseModels.TransformedTeam) []string {
	teams := addTeamsInSlice(modelTeams)
	checkTeamNumber(teams)
	return teams
}

func addTeamsInSlice(modelTeams []databaseModels.TransformedTeam) []string {
	var teams []string

	for i := 0; i < len(modelTeams); i++ {
		teams = append(teams, modelTeams[i].TeamName)
	}

	return teams
}

func checkTeamNumber(teams []string) {
	if isTeamsOdd(len(teams)) {
		teams = append(teams, "absent")
	}
}

func isTeamsOdd(teamNumber int) bool {
	return teamNumber != 0
}

func createSchedule(teams []string, robin int, id uint) []models.Match {
	var matches []models.Match

	for i := 0; i < (len(teams)-1)*robin; i++ {
		firstHalf, secondHalf := divideTeamSliceInHalf(i, teams)
		matches = makePairs(matches, firstHalf, secondHalf, i, id)
		activateRobin(teams)
	}

	return matches
}

func divideTeamSliceInHalf(round int, teams []string) ([]string, []string) {
	var firstHalf []string
	var secondHalf []string

	if isRoundEven(round) {
		firstHalf = teams[:len(teams)/2]
		secondHalf = teams[len(teams)/2:]
	} else {
		firstHalf = teams[len(teams)/2:]
		secondHalf = teams[:len(teams)/2]
	}

	return firstHalf, secondHalf
}

func isRoundEven(round int) bool {
	return round%2 == 0
}

func makePairs(matches []models.Match, firstHalf []string, secondHalf []string, round int, id uint) []models.Match {
	var pairs []models.Pair

	for k := 0; k < len(firstHalf); k++ {
		pairs = append(pairs, models.Pair{
			Number:    round*len(firstHalf) + k + 1,
			HomeTeam:  firstHalf[k],
			HomeScore: 0,
			AwayTeam:  secondHalf[len(secondHalf)-1-k],
			AwayScore: 0,
		})
	}

	matches = append(matches, models.Match{
		TournamentId: id,
		Round:        strconv.Itoa(round + 1),
		Pair:         pairs,
	})

	return matches
}

func activateRobin(teams []string) {
	temp := teams[len(teams)-1]
	for j := len(teams) - 1; j > 1; j-- {
		teams[j] = teams[j-1]
	}
	teams[1] = temp
}

func madeJSON(matches []models.Match) []byte {
	matchesJSON, err := json.Marshal(matches)
	if err != nil {
		panic(err)
	}

	return matchesJSON
}
