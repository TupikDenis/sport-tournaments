package token

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"sport-tournaments/pkg/models/databaseModels"
	"time"
)

var signingKey string = "grkjk#4#35FSFJLja#4353KSFjH"

func GenerateToken(user databaseModels.TransformedUser) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &databaseModels.Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		Id:       user.Id,
		Username: user.Username,
		Role:     user.Role,
	})

	return token.SignedString([]byte(signingKey))
}

func ParseToken(accessToken string) (databaseModels.TransformedUser, error) {
	token, err := jwt.ParseWithClaims(accessToken, &databaseModels.Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(signingKey), nil
	})

	var user databaseModels.TransformedUser

	if err != nil {
		return user, err
	}

	if claims, ok := token.Claims.(*databaseModels.Claims); ok && token.Valid {
		user = databaseModels.TransformedUser{
			Id:       claims.Id,
			Username: claims.Username,
			Role:     claims.Role,
		}
		return user, nil
	}

	return user, errors.New("invalid auth token")
}
