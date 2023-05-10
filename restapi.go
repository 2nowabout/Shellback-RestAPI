package main

import (
	"Shellback.nl/Restapi/controllers"
	"Shellback.nl/Restapi/middlewares"
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
	aliveRepo := controllers.NewAliveRepo()
	loginRepo := controllers.NewLoginRepo()

	router.POST("/register", loginRepo.Register)
	router.POST("/login", loginRepo.Login)

	router.POST("/Alive", aliveRepo.UpdateAlive)
	router.GET("/updateActive/:ip", companyRepo.UpdateLastActive)
	router.POST("/addNotification", notificationRepo.CreateNotification)

	secured := router.Group("").Use(middlewares.JwtAuthMiddleware())
	secured.GET("/getNotifications/:ip", notificationRepo.GetNotifications)
	secured.GET("/getCompanies", companyRepo.GetCompanys)
	secured.GET("/getCompany/:id", companyRepo.GetCompany)
	secured.GET("/getTypes", typeRepo.GetTypes)
	secured.GET("/getAllAlive", aliveRepo.GetAlives)
	secured.GET("/getAlive/:id", aliveRepo.GetAlive)
	secured.POST("/addCompany", companyRepo.CreateCompany)
	secured.DELETE("/deleteCompany/:id", companyRepo.DeleteCompany)
	secured.DELETE("/deleteNotifications/:ip", notificationRepo.DeleteNotifications)
	router.Run("localhost:8002")
}
