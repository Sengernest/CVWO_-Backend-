package main

import "my-backend/models"

func GetUsers() ([]models.User, error) {
	var users []models.User
	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil

}
