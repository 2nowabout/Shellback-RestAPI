package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"Shellback.nl/Restapi/database"
	"Shellback.nl/Restapi/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LoginRepo struct {
	Db *gorm.DB
}

func NewLoginRepo() *LoginRepo {
	db := database.InitDb()
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.User{})
	return &LoginRepo{Db: db}
}

func (repository *LoginRepo) Login(c *gin.Context) {

	var input models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.User{}

	u.Username = input.Username
	u.Password = input.Password
	token, err := models.AuthenticateUser(repository.Db, u.Username, u.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})

}

func (repository *LoginRepo) Register(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u := models.User{}
	u.Username = input.Username
	u.Email = input.Email
	u.Password = input.Password

	err := models.CreateUser(repository.Db, &u)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "registration success"})
}

func (repository *LoginRepo) GetUserByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var User models.User
	User, err := models.GetUserByID(repository.Db, uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, User)
}
