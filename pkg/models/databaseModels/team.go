package databaseModels

import "gorm.io/gorm"

type Team struct {
	gorm.Model
	TournamentId uint   `json:"tournament_id"`
	TeamName     string `json:"team_name"`
	Win          uint   `json:"win"`
	Draw         uint   `json:"draw"`
	Lose         uint   `json:"lose"`
	Score        uint   `json:"score"`
	Against      uint   `json:"against"`
}

type TransformedTeam struct {
	Id           uint   `json:"id"`
	TournamentId uint   `json:"tournament_id"`
	TeamName     string `json:"team_name"`
	Win          uint   `json:"win"`
	Draw         uint   `json:"draw"`
	Lose         uint   `json:"lose"`
	Score        uint   `json:"score"`
	Against      uint   `json:"against"`
}
