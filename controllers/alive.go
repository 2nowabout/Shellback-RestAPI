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

type AliveRepo struct {
	Db *gorm.DB
}

func NewAliveRepo() *AliveRepo {
	db := database.InitDb()
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.Alive{})
	return &AliveRepo{Db: db}
}

// create Alive
func (repository *AliveRepo) CreateAlive(c *gin.Context) {
	var Alive models.Alive
	c.BindJSON(&Alive)
	err := models.CreateAlive(repository.Db, &Alive)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, Alive)
	return
}

// get Alives
func (repository *AliveRepo) GetAlives(c *gin.Context) {
	var Alive []models.Alive
	err := models.GetAlives(repository.Db, &Alive)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, Alive)
	return
}

// get Alive by id
func (repository *AliveRepo) GetAlive(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var Alive models.Alive
	err := models.GetAlive(repository.Db, &Alive, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, Alive)
	return
}

// delete Alive
func (repository *AliveRepo) DeleteAlive(c *gin.Context) {
	var Alive models.Alive
	id, _ := strconv.Atoi(c.Param("id"))
	err := models.DeleteAlive(repository.Db, &Alive, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Alive deleted successfully"})
	return
}

func (repository *AliveRepo) UpdateAlive(c *gin.Context) {
	ip := c.Param("ip")
	var Alive models.Alive
	err := models.UpdateActive(repository.Db, ip)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, Alive)
}
