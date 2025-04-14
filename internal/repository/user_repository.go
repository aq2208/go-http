package repository

import (
	"go-database/internal/model"
	"go-database/internal/database"
)

func GetAllUsers() ([]model.User, error) {
	rows, err := database.DB.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var u model.User
		if err := rows.Scan(&u.UserId, &u.Name, &u.Email, &u.Address, &u.CreatedAt); err != nil {
			return nil, err
		}

		users = append(users, u)
	}

	return users, nil
}