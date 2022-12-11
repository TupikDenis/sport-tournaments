package user

import (
	"sport-tournaments/pkg/handlers"
	"sport-tournaments/pkg/models/databaseModels"
	"sport-tournaments/pkg/services/token"
)

func Login(username string, password string) string {
	var users []databaseModels.User
	var _users []databaseModels.TransformedUser
	var user databaseModels.TransformedUser

	db := handlers.Database()
	db.Find(&users)

	for _, item := range users {
		if item.Username == username && item.Password == password {
			_users = append(
				_users, databaseModels.TransformedUser{
					Id:       item.ID,
					Username: item.Username,
					Role:     item.Role,
				})
		}
	}

	var tokenStr string

	if len(_users) == 1 {
		user = _users[0]
		tokenTmp, err := token.GenerateToken(user)

		if err != nil {
			panic(err)
		} else {
			tokenStr = tokenTmp
		}
	}

	return tokenStr
}

func GetUserPassword(id uint) string {
	var users []databaseModels.User
	var password string

	db := handlers.Database()
	db.Find(&users)

	for _, item := range users {
		if item.ID == id {
			password = item.Password
		}
	}

	return password
}
