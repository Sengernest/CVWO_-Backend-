package main

import (
	"log"
	"my-backend/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	var err error
	// Open the database connection
	db, err = gorm.Open(sqlite.Open("database.db"), &gorm.Config{TranslateError: true})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Migrate models
	if err := db.AutoMigrate(&models.User{}, &models.Comment{}, &models.Thread{}); err != nil {
		log.Fatalf("Failed to migrate models: %v", err)
	}

	// Create Gin router
	router := gin.Default()

	// Routes
	router.GET("/threads", func(ctx *gin.Context) {
		threads, err := GetThreads()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, threads)
	})

	router.GET("/threads/:id/comments", func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		comments, err := GetComments(uint(id))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, comments)
	})

	router.GET("/users", func(ctx *gin.Context) {
		users, err := GetUsers()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, users)
	})

	router.POST("/threads", func(ctx *gin.Context) {
		var requestBody models.Thread
		if err := ctx.ShouldBindJSON(&requestBody); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		thread, err := CreateThread(&requestBody)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, thread)
	})

	router.PATCH("/threads/:id", func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		var requestBody models.Thread
		if err := ctx.ShouldBindJSON(&requestBody); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		thread, err := UpdateThread(uint(id), requestBody.Title, requestBody.Content)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, thread)
	})

	router.DELETE("/threads/:id", func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if err := DeleteThread(uint(id)); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.Status(http.StatusNoContent)
	})

	// Run server
	router.Run(":8080")
}

// Helper functions (the same as before)
func GetThreads() ([]models.Thread, error) {
	var threads []models.Thread
	if err := db.Find(&threads).Error; err != nil {
		return nil, err
	}
	return threads, nil
}

func CreateThread(thread *models.Thread) (*models.Thread, error) {
	if err := db.Create(thread).Error; err != nil {
		return nil, err
	}
	return thread, nil
}

func UpdateThread(ID uint, Title string, Content string) (*models.Thread, error) {
	var thread models.Thread
	if err := db.First(&thread, ID).Error; err != nil {
		return nil, err
	}

	if err := db.Model(&thread).Updates(models.Thread{Title: Title, Content: Content}).Error; err != nil {
		return nil, err
	}

	return &thread, nil
}

func DeleteThread(ID uint) error {
	result := db.Delete(&models.Thread{}, ID)
	return result.Error
}

func GetComments(ThreadID uint) ([]models.Comment, error) {
	var comments []models.Comment
	if err := db.Where("thread_id = ?", ThreadID).Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}

func GetUsers() ([]models.User, error) {
	var users []models.User
	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
