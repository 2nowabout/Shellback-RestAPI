package controllers

import (
	"fmt"
	"net/http"

	"Shellback.nl/Restapi/database"
	"Shellback.nl/Restapi/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type NotificationRepo struct {
	Db *gorm.DB
}

func NewNotificationRepo() *NotificationRepo {
	db := database.InitDb()
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.Notification{})
	return &NotificationRepo{Db: db}
}

// create Notification
func (repository *NotificationRepo) CreateNotification(c *gin.Context) {
	var Notification models.Notification
	c.BindJSON(&Notification)
	err := models.CreateNotification(repository.Db, &Notification)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, Notification)
}

// get Notifications
func (repository *NotificationRepo) GetNotifications(c *gin.Context) {
	fmt.Println("TRYING TO FIND IT")
	IpAdress := c.Param("ip")
	var Notification []models.Notification
	err := models.GetNotifications(repository.Db, &Notification, IpAdress)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, Notification)
}

// delete Notification
func (repository *NotificationRepo) DeleteNotifications(c *gin.Context) {
	var Notification []models.Notification
	id := c.Param("ip")
	err := models.DeleteNotifications(repository.Db, &Notification, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Notification deleted successfully"})
}
