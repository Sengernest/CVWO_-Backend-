package main

import "my-backend/models"

func GetComments(ThreadID uint) ([]models.Comment, error) {
	var comments []models.Comment
	if err := db.Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}

