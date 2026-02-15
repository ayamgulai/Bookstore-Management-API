package services

import (
	"errors"

	"bookstore-management-api/repositories"
	"bookstore-management-api/utils"
)

func Login(username, password string) (string, error) {
	user, err := repositories.GetUserByUsername(username)
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	// compare hashed password
	if !utils.CheckPasswordHash(password, user.Password) {
		return "", errors.New("invalid username or password")
	}

	return utils.GenerateToken(user.ID, user.Username)
}
