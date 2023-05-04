package main

import (
	"Shellback.nl/Restapi/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	requestHandler()
}

func requestHandler() {
	router := gin.Default()
	companyRepo := controllers.NewCompanyRepo()
	notificationRepo := controllers.NewNotificationRepo()
	typeRepo := controllers.NewTypeRepo()

	router.GET("/getNotifications/:ip", notificationRepo.GetNotifications)
	router.GET("/getCompanies", companyRepo.GetCompanys)
	router.GET("/getCompany/:id", companyRepo.GetCompany)
	router.GET("/getTypes", typeRepo.GetTypes)
	router.POST("/addNotification", notificationRepo.CreateNotification)
	router.POST("/addCompany", companyRepo.CreateCompany)
	router.POST("/alive")
	router.GET("/updateActive/:ip", companyRepo.UpdateLastActive)
	router.DELETE("/deleteCompany/:id", companyRepo.DeleteCompany)
	router.DELETE("/deleteNotifications/:ip", notificationRepo.DeleteNotifications)
	router.Run("localhost:8002")
}
