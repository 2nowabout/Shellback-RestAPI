package main

import (
	"Shellback.nl/Restapi/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	//requestHandler()
	getDotEnv()
}

func requestHandler() {
	router := gin.Default()
	companyRepo := controllers.NewCompanyRepo()
	notificationRepo := controllers.NewNotificationRepo()
	typeRepo := controllers.NewTypeRepo()

	router.GET("/getNotifications/:ip", notificationRepo.GetNotifications)
	router.GET("/getCompanys", companyRepo.GetCompanys)
	router.GET("/getCompany/:id", companyRepo.GetCompany)
	router.GET("/getTypes", typeRepo.GetTypes)
	router.POST("/addNotification", notificationRepo.CreateNotification)
	router.POST("/addCompany", companyRepo.CreateCompany)
	router.DELETE("/deleteCompany/:id", companyRepo.DeleteCompany)
	router.DELETE("/deleteNotifications/:ip", notificationRepo.DeleteNotifications)

	router.Run("localhost:8000")
}
