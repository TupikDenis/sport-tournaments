package databaseModels

import (
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type TransformedUser struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

type Claims struct {
	jwt.StandardClaims
	Id       uint   `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
}
