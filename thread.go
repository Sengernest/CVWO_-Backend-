package main

import "my-backend/models"

func GetThreads() ([]models.Thread, error) {
	var threads []models.Thread
	if err := db.Find(&threads).Error; err != nil {
		return nil, err
	}
	return threads, nil

}

func CreateThread(thread  *models.Thread) (*models.Thread, error)  {
	if error := db.Create(thread).Error; error != nil {
		return nil, error
	} 
	return thread, nil
}

func UpdateThread(ID uint, Title string, Content string) (*models.Thread, error) {
	var thread models.Thread
	if error := db.First(&thread, ID).Error; error != nil {
		return nil, error
	}
	
	if error := db.Model(&thread).Updates(models.Thread{Title: Title, Content: Content}).Error; error != nil {
		return nil, error
	}

	return &thread, nil 
} 

func DeleteThread(ID uint) (error) {
	result := db.Delete(&models.Thread{}, ID)
	if result.Error != nil {
		return result.Error
	}

	return nil 
}
