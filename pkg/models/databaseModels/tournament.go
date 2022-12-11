package databaseModels

import "github.com/jinzhu/gorm"

type Tournament struct {
	gorm.Model
	UsernameId  uint   `json:"username"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Sport       string `json:"sport"`
	System      string `json:"system"`
}

type TransformedTournament struct {
	Id          uint   `json:"id"`
	UsernameId  uint   `json:"username"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Sport       string `json:"sport"`
	System      string `json:"system"`
}
