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

		//* scan and assign each row's column values to struct fields (ORDER MATTERS)
		if err := rows.Scan(&u.UserId, &u.Name, &u.Email, &u.Address, &u.CreatedAt); err != nil {
			return nil, err
		}

		users = append(users, u)
	}

	return users, nil
}

func GetUserById(userId int) (*model.User, error) {
	row, err := database.DB.Query("SELECT * FROM users WHERE user_id = ?", userId)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	var u model.User

	row.Next()
	if err := row.Scan(&u.UserId, &u.Name, &u.Email, &u.Address, &u.CreatedAt); err != nil {
		return nil, err
	}

	return &u, nil
}