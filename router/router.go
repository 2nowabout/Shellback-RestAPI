package router

import (
	"Shellback.nl/Restapi/controllers"
	"Shellback.nl/Restapi/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func RequestHandler() *gin.Engine {
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

	router.POST("/login", loginRepo.Login)

	router.POST("/Alive", aliveRepo.UpdateAlive)
	router.GET("/updateActive/:ip", companyRepo.UpdateLastActive)
	router.POST("/addNotification", notificationRepo.CreateNotification)
	router.POST("/addCompany", companyRepo.CreateCompany)

	secured := router.Group("").Use(middlewares.JwtAuthMiddleware())
	secured.POST("/register", loginRepo.Register)
	secured.POST("/changepassword", loginRepo.UpdatePassword)

	secured.GET("/getNotifications/:ip", notificationRepo.GetNotifications)

	secured.POST("/updateCompanyName", companyRepo.UpdateCompanyName)
	secured.GET("/getCompanies", companyRepo.GetCompanys)
	secured.GET("/getCompany/:id", companyRepo.GetCompany)
	secured.GET("/getCompany/notificationAmount/:ip", notificationRepo.GetAmountNotifications)

	secured.GET("/getTypes", typeRepo.GetTypes)

	secured.GET("/getAllAlive", aliveRepo.GetAlives)
	secured.GET("/getAlive/:id", aliveRepo.GetAlive)

	secured.DELETE("/deleteCompany/:id", companyRepo.DeleteCompany)
	secured.DELETE("/deleteNotifications/:ip", notificationRepo.DeleteNotifications)
	return router
}
