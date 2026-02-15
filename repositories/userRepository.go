package repositories

import (
	"bookstore-management-api/configs"
	"bookstore-management-api/models"
)

func GetUserByUsername(username string) (*models.User, error) {
	row := configs.DB.QueryRow(`
		SELECT id, username, password, created_at
		FROM users WHERE username = $1
	`, username)

	var user models.User
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
