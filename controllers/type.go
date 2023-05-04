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

type TypeRepo struct {
	Db *gorm.DB
}

func NewTypeRepo() *TypeRepo {
	db := database.InitDb()
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.Type{})
	return &TypeRepo{Db: db}
}

// create Type
func (repository *TypeRepo) CreateType(c *gin.Context) {
	var Type models.Type
	c.BindJSON(&Type)
	err := models.CreateType(repository.Db, &Type)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, Type)
}

// get Types
func (repository *TypeRepo) GetTypes(c *gin.Context) {
	var Type []models.Type
	err := models.GetTypes(repository.Db, &Type)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, Type)
}

// get Type by id
func (repository *TypeRepo) GetType(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var Type models.Type
	err := models.GetType(repository.Db, &Type, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, Type)
}

// update Type
func (repository *TypeRepo) UpdateType(c *gin.Context) {
	var Type models.Type
	id, _ := strconv.Atoi(c.Param("id"))
	err := models.GetType(repository.Db, &Type, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.BindJSON(&Type)
	err = models.UpdateType(repository.Db, &Type)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, Type)
}

// delete Type
func (repository *TypeRepo) DeleteType(c *gin.Context) {
	var Type models.Type
	id, _ := strconv.Atoi(c.Param("id"))
	err := models.DeleteType(repository.Db, &Type, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Type deleted successfully"})
}
