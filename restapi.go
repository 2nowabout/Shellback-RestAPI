package main

import (
	"Shellback.nl/Restapi/controllers"
	"Shellback.nl/Restapi/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	requestHandler()
}

func requestHandler() {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowMethods = []string{"GET", "POST", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Content-Type", "Authorization"}
	config.AllowAllOrigins = true
	router.Use(cors.New(config))
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
	secured.GET("/getCompany/notificationAmount/:ip", notificationRepo.GetAmountNotifications)
	secured.GET("/getTypes", typeRepo.GetTypes)
	secured.GET("/getAllAlive", aliveRepo.GetAlives)
	secured.GET("/getAlive/:id", aliveRepo.GetAlive)
	secured.POST("/addCompany", companyRepo.CreateCompany)
	secured.DELETE("/deleteCompany/:id", companyRepo.DeleteCompany)
	secured.DELETE("/deleteNotifications/:ip", notificationRepo.DeleteNotifications)
	router.Run("localhost:8002")
}
