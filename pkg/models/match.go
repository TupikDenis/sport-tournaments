package models

import "gorm.io/gorm"

type Pair struct {
	Number    int    `json:"number"`
	HomeTeam  string `json:"home_team"`
	HomeScore int    `json:"home_score"`
	AwayTeam  string `json:"away_team"`
	AwayScore int    `json:"away_score"`
}

type Match struct {
	TournamentId uint   `json:"tournament_id"`
	Round        string `json:"round"`
	Pair         []Pair `json:"pairs"`
}

type FullMatch struct {
	gorm.Model
	TournamentId uint
	Round        string
	Robin        int
	Number       int
	HomeTeam     string
	HomeScore    int
	AwayTeam     string
	AwayScore    int
}

type TransformedFullMatch struct {
	Id           uint
	TournamentId uint
	Round        string
	Robin        int
	Number       int
	HomeTeam     string
	HomeScore    int
	AwayTeam     string
	AwayScore    int
}
